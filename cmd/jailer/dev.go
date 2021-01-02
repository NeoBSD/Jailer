package main

import (
	"github.com/spf13/cobra"
)

// devCmd represents the `dev` sub command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Subcommand for development only",
	RunE:  RunDevCommand,
}

// RunDevCommand executes the dev subcommand.
func RunDevCommand(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	rootCmd.AddCommand(devCmd)
}
