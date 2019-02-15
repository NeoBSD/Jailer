package cmd

import (
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// storageCmd represents the config sub command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manages container & image storage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Start a process:
		c := exec.Command("zfs", "list")
		if err := c.Run(); err != nil {
			logrus.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(storageCmd)

	// Here you will define your flags and configuration settings.

}
