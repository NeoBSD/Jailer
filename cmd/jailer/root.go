package main

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	tabWriterPadding = 3
)

// ErrSilent ...
var ErrSilent = errors.New("ErrSilent")

var rootCmd = &cobra.Command{
	Use:           "jailer",
	Short:         "jailer",
	Long:          `jailer https://github.com/NeoBSD/jailer`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Println(err)
		cmd.Println(cmd.UsageString())
		return ErrSilent
	})

	// Viper config
	cobra.OnInitialize(initConfig)

	// Flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().Bool("json", false, "JSON output")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))

}

func initConfig() {
	// // Will be uppercased automatically
	// viper.SetEnvPrefix("jailer")
	// viper.BindEnv("path")

	// os.Setenv("JAILER_PATH", "13") // typically done outside of the app
}
