package cmd

import (
	"github.com/discoriver/building-tool-suite/internal/terminator"
	"github.com/discoriver/building-tool-suite/pkg/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	// Config keys. I use "core" for program-wide config items.
	DebugConfigKey              = "core.Debug"

	// Custom config?
	configFileFlag string

	// Hold debug log value so we can assign it properly
	debugLogs = false

	rootCmd = &cobra.Command{
		Use:   "bts",
		Short: "building-tool-suite (bts) is a sample program outlining the design of internals tools.",
		Long:  "building-tool-suite (bts) is a sample program outlining the design of internals tools into an easy-to-understand and maintainable package.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Get our config before anything else.
	cobra.OnInitialize(initConfig, initLogger)

	rootCmd.PersistentFlags().StringVar(&configFileFlag, "config", "", "Custom bts config")
	rootCmd.PersistentFlags().BoolVar(&debugLogs, "debug", false, "Enable debug logging")

	viper.BindPFlag(DebugConfigKey, rootCmd.PersistentFlags().Lookup("debug"))

	// Set config defaults
	viper.SetDefault(terminator.DBTableHeaderForegroundColourConfigKey, "white")
	viper.SetDefault(terminator.DBTableHeaderBackgroundColourConfigKey, "")
}

func initLogger() {
	log.Verbose = viper.GetBool(DebugConfigKey)
}

func initConfig() {
	defaultConfigName := ".btsconf"
	defaultConfigType := "yaml"

	var configHome string

	if configFileFlag != "" {
		// User config file from the flag.
		viper.SetConfigFile(configFileFlag)
	} else {
		// Find log file in home directory (default location)
		var err error
		configHome, err = homedir.Dir()
		if err != nil {
			log.Error("%s", err)
		}

		viper.AddConfigPath(configHome)
		viper.SetConfigName(defaultConfigName)
		viper.SetConfigType(defaultConfigType)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug("CONFIG:: %s", "Config file not found, using defaults")
		} else {
			// We don't want to use defaults if user is trying to use a custom config, ideally.
			log.Error("CONFIG:: %s", "Config file found, but errored: %s", err)
		}
	}

	// Check config for correct permissions. Required due to the sensitive nature of it's contents.
	file := viper.ConfigFileUsed()
	f, err := os.Stat(file)
	if err == nil && f.Mode().Perm() != 0600 {
		log.Error("CONFIG:: Config file %s has invalid permissions. Run \"chmod 0600 %s\" to correct.", file, file)
	}
}
