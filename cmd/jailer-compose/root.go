package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "jailer-compose",
	Short: "jailer-compose",
	Long:  `jailer-compose https://github.com/NeoBSD/jailer`,
	Run:   nil,
}

func init() {
	rootCmd.SetIn(os.Stdin)
	rootCmd.SetOut(os.Stdout)

	fileHelp := "Specify an alternate compose file"
	rootCmd.PersistentFlags().StringP("file", "f", "jailer-compose.yml", fileHelp)

	projectHelp := "Specify an alternate project name (default: directory name)"
	rootCmd.PersistentFlags().StringP("project-name", "p", "", projectHelp)

	projectDirectoryHelp := "Specify an alternate working directory (default: the path of the compose file)"
	rootCmd.PersistentFlags().String("project-directory", "", projectDirectoryHelp)

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().Bool("json", false, "JSON output")

	viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
	viper.BindPFlag("project-name", rootCmd.PersistentFlags().Lookup("project-name"))
	viper.BindPFlag("project-directory", rootCmd.PersistentFlags().Lookup("project-directory"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))

}

func initConfig() {
}
