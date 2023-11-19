/*
Copyright © 2023 saltfishpr <526191197@qq.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"time"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/go-audio/wav"
	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/vansante/go-ffprobe.v2"
)

type Model struct {
	name     string
	url      string
	checksum string
}

var models = map[string]Model{
	"tiny": {
		name:     "ggml-tiny.bin",
		url:      "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-tiny.bin",
		checksum: "be07e048e1e599ad46341c8d2a135645097a538221678b7acdd1b1919c6e1b21",
	},
	"base": {
		name:     "ggml-base.bin",
		url:      "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-base.bin",
		checksum: "60ed5bc3dd14eea856493d334349b405782ddcaf0028d4b5df4088345fba2efe",
	},
	"small": {
		name:     "ggml-small.bin",
		url:      "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-small.bin",
		checksum: "1be3a9b2063867b937e64e2ec7483364a79917e157fa98c5d94b5c1fffea987b",
	},
	"medium": {
		name:     "ggml-medium.bin",
		url:      "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-medium.bin",
		checksum: "6c14d5adee5f86394037b4e4e8b59f1673b6cee10e3cf0b11bbdbee79c156208",
	},
	"large": {
		name:     "ggml-large.bin",
		url:      "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-large.bin",
		checksum: "", // TODO
	},
	"largev3": {
		name:     "ggml-large.bin",
		url:      "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-large-v3.bin",
		checksum: "",
	},
}

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "go-whisper",
	Short:   "go-whisper can transcribe audio to text",
	Long:    `go-whisper can transcribe audio to text`,
	Version: "v0.1.1",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := contextForSignal(os.Interrupt)

		var input string
		switch len(args) {
		case 0:
			return errors.New("input file is required")
		case 1:
			input = args[0]
		default:
			input = args[0]
			fmt.Println("Warning: only the first input file is used")
		}

		if _, err := os.Stat(input); err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("input file %s not exists", input)
			}
			return err
		}

		format, err := ffprobe.ProbeURL(ctx, input)
		if err != nil {
			return fmt.Errorf("probe audio file failed: %w", err)
		}
		if format.Format.FormatName != "wav" &&
			format.FirstAudioStream().SampleRate != "16000" &&
			format.FirstAudioStream().Channels != 1 {
			_input := path.Join(os.TempDir(), "go-whisper-tmp.wav")
			cmd := exec.CommandContext(ctx, "ffmpeg", "-y", "-i", input, "-ac", "1", "-ar", "16000", _input)
			if verbose {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
			}
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("convert audio format failed: %w", err)
			}
			input = _input
		}

		model, err := cmd.Flags().GetString("model")
		if err != nil {
			return err
		}
		if !lo.Contains(lo.Keys(models), model) {
			return fmt.Errorf("model %s is not supported", model)
		}

		cacheDir, err := os.UserCacheDir()
		if err != nil {
			return err
		}
		modelDir := path.Join(cacheDir, "go-whisper/models")
		if err := os.MkdirAll(modelDir, 0o755); err != nil {
			return err
		}
		modelPath := path.Join(modelDir, models[model].name)
		if _, err := os.Stat(modelPath); err != nil {
			if os.IsNotExist(err) {
				if err := downloadModel(ctx, modelDir, model); err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			checksum, err := sha256sum(modelPath)
			if err != nil {
				return err
			}
			// checksum mismatch, re-download
			if checksum != models[model].checksum {
				os.Remove(modelPath)
				if err := downloadModel(ctx, modelDir, model); err != nil {
					return err
				}
			}
		}

		checksum, err := sha256sum(modelPath)
		if err != nil {
			return err
		}
		if checksum != models[model].checksum {
			return fmt.Errorf("model checksum mismatch: %s != %s", checksum, models[model].checksum)
		}

		// whisper model
		wm, err := whisper.New(modelPath)
		if err != nil {
			return fmt.Errorf("create whisper model failed: %w", err)
		}
		defer wm.Close()

		// whisper context
		wc, err := wm.NewContext()
		if err != nil {
			return fmt.Errorf("create whisper context failed: %w", err)
		}

		lang, err := cmd.Flags().GetString("language")
		if err != nil {
			return err
		}
		if err := wc.SetLanguage(lang); err != nil {
			return fmt.Errorf("set language failed: %w", err)
		}

		fh, err := os.Open(input)
		if err != nil {
			return err
		}
		defer fh.Close()

		var data []float32
		dec := wav.NewDecoder(fh)
		if buf, err := dec.FullPCMBuffer(); err != nil {
			return err
		} else if dec.SampleRate != whisper.SampleRate {
			return fmt.Errorf("unsupported sample rate: %d", dec.SampleRate)
		} else if dec.NumChans != 1 {
			return fmt.Errorf("unsupported number of channels: %d", dec.NumChans)
		} else {
			data = buf.AsFloat32Buffer().Data
		}

		var cb whisper.SegmentCallback = func(s whisper.Segment) {
			fmt.Printf("%02d [%4s->%4s] ", s.Num, s.Start.Truncate(time.Millisecond), s.End.Truncate(time.Millisecond))
			fmt.Println(s.Text)
		}

		wc.ResetTimings()
		if err := wc.Process(data, cb, nil); err != nil {
			return err
		}
		wc.PrintTimings()

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-whisper.yaml)")

	rootCmd.Flags().StringP("model", "m", "base", "whisper model to use\neg. tiny, base, small, medium or large")
	rootCmd.Flags().StringP("language", "l", "auto", "language to use for speech recognition\neg. Chinese, Japanese, English, French, etc.")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show verbose output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-whisper" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-whisper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// contextForSignal returns a context object which is cancelled when a signal
// is received. It returns nil if no signal parameter is provided
func contextForSignal(signals ...os.Signal) context.Context {
	if len(signals) == 0 {
		return nil
	}

	ch := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	// Send message on channel when signal received
	signal.Notify(ch, signals...)

	// When any signal received, call cancel
	go func() {
		<-ch
		cancel()
	}()

	// Return success
	return ctx
}

func sha256sum(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func downloadModel(ctx context.Context, folder string, model string) error {
	modelPath := path.Join(folder, models[model].name)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, models[model].url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(modelPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("create temp file error: %w", err)
	}
	defer f.Close()

	fmt.Printf("downloading model %s to %s \n", model, modelPath)
	bar := progressbar.DefaultBytes(resp.ContentLength, "downloading")
	if _, err := io.Copy(io.MultiWriter(f, bar), resp.Body); err != nil {
		return fmt.Errorf("download model %s failed: %w", model, err)
	}
	return nil
}
