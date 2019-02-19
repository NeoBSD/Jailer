package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// fetchCmd represents the config sub command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch base from FreeBSD mirror",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Get OS version
		c := exec.Command("freebsd-version")
		stdout, err := c.Output()
		if err != nil {
			return fmt.Errorf("Error while executing external command: %v", err.Error())
		}

		fmt.Print(string(stdout))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

}