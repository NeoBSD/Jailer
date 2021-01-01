package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/NeoBSD/jailer"
)

// versionCmd represents the version sub command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current version",
	RunE:  RunVersionCommand,
}

// RunVersionCommand prints the current jailer version
func RunVersionCommand(cmd *cobra.Command, args []string) error {

	if viper.Get("verbose") == true {
		fmt.Printf("Version: %s\n", jailer.Version)
		fmt.Printf("Commit: %s\n", jailer.BuildCommit)
		fmt.Printf("Date: %s\n", jailer.BuildDate)
		fmt.Printf("Build on: %s\n", jailer.BuildOS)
		return nil
	}

	fmt.Printf("%s-%s\n", jailer.Version, jailer.BuildCommit)
	return nil
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
