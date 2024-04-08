package main

import (
	"os"

	"braces.dev/errtrace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	Rendezvous  string `mapstructure:"rendezvous"`
	Protocol    string `mapstructure:"protocol"`
	BackendHost string `mapstructure:"host"`
	BackendPort int    `mapstructure:"port"`
	ListenHost  string `mapstructure:"p2p_host"`
	ListenPort  int    `mapstructure:"p2p_port"`
}

// initConfig reads in config file and ENV variables if set.
func initConfig() (*config, error) {
	var cfgFile string
	pflag.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chatp2p.yaml)")
	pflag.String("host", "0.0.0.0", "The bootstrap node host listen address")
	pflag.Int("port", 4001, "node listen port")
	pflag.Parse()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".chatp2p" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".chatp2p")
	}

	viper.SetDefault("host", "127.0.0.1")
	viper.SetDefault("port", 4001)
	viper.SetDefault("p2p_host", "0.0.0.0")
	viper.SetDefault("p2p_port", 0)

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, errtrace.Wrap(err)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return nil, errtrace.Wrap(err)
	}

	cfg := new(config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errtrace.Wrap(err)
	}
	return cfg, nil
}
