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
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"

	"learning/internal"

	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:     "convert [--timezone name] [--reload] 'yyyy-mm-dd hh:mm:ss'",
	Short:   "Convert time string to timestamp",
	Example: "convert --timezone Asia/Shanghai '2021-01-01 12:00:00'",
	Args:    cobra.ExactArgs(1),

	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := internal.ContextForSignal(os.Interrupt)

		loader := internal.NewLoader("data")
		tzs, err := loader.Load(ctx, viper.GetBool("reload"))
		if err != nil {
			return fmt.Errorf("load timezones error: %w", err)
		}

		name := viper.GetString("timezone")
		if name == "" {
			return fmt.Errorf("input timezone is empty")
		}

		timeStr := args[0]
		if _, err := time.ParseInLocation(time.DateTime, timeStr, time.UTC); err != nil {
			return fmt.Errorf("invalid time string '%s'", timeStr)
		}

		slog.Info(fmt.Sprintf("convert '%s' to timestamp in %s", timeStr, name))

		var res []string
		for _, tz := range tzs {
			if tz.Name != name {
				continue
			}
			if !tz.Contains(timeStr) {
				continue
			}
			ts := lo.Must(tz.Convert(timeStr))
			res = append(res, cast.ToString(ts))
		}
		if len(res) == 0 {
			// 如果没有解析出来, 则尝试用 golang 内置的时区解析器
			loc, err := time.LoadLocation(name)
			if err != nil {
				return fmt.Errorf("load timezone '%s' error: %w", name, err)
			}
			ts := lo.Must(time.ParseInLocation(time.DateTime, timeStr, loc)).Unix()
			res = append(res, cast.ToString(ts))
		}
		fmt.Println(strings.Join(res, ","))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	const defaultTimezone = "Asia/Shanghai"
	convertCmd.Flags().String("timezone", "", "The timezone to use.")
	checkBindError(viper.BindPFlag("timezone", convertCmd.Flags().Lookup("timezone")))
	if runtime.GOOS == "windows" {
		viper.SetDefault("timezone", defaultTimezone)
	} else {
		name, err := getLinuxTimezone()
		if err != nil {
			viper.SetDefault("timezone", defaultTimezone)
		} else {
			viper.SetDefault("timezone", name)
		}
	}

	convertCmd.Flags().Bool("reload", false, "Reload data from timezone db.")
	checkBindError(viper.BindPFlag("reload", convertCmd.Flags().Lookup("reload")))
}

func getLinuxTimezone() (string, error) {
	path, err := os.Readlink("/etc/localtime")
	if err != nil {
		return "", err
	}
	// remove prefix "/usr/share/zoneinfo/" from p
	return path[20:], nil
}
