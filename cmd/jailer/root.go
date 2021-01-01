package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	tabWriterPadding = 3
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "jailer",
	Short: "jailer",
	Long:  `jailer https://github.com/NeoBSD/jailer`,
	Run:   nil,
}

func init() {
	// Viper config
	cobra.OnInitialize(initConfig)

	// Flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $PWD/jailer.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find pwd
		dir, err := os.Getwd()
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}

		// Search config in home directory with name "jailer" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("jailer")
		viper.AutomaticEnv() // read in environment variables that match

	}

	// Set log level
	// TODO: set via jailer.yml
	if viper.Get("verbose") == true {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
