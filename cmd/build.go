package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the config sub command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a jailer container",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

}
