package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the config sub command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a jailer container",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

}
