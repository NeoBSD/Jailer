package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tobiashienzsch/jailer/filesystem"
)

// initCmd represents the config sub command
var initCmd = &cobra.Command{
	Use:   "init [zpool]",
	Short: "Init for jailer. Creates zfs datasets",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		// Check zpool passed
		// Create dataset
		jailerRoot := fmt.Sprintf("%s/jailer", args[0])
		filesystem.NewDataset(jailerRoot)

		// Write initial config
		// Enable jails in rc.conf

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

}
