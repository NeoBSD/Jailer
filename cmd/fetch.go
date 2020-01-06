package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// fetchCmd represents the config sub command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch base from FreeBSD mirror",
	Run:   RunFetchCommand,
}

// RunFetchCommand ...
func RunFetchCommand(cmd *cobra.Command, args []string) {

	// Get OS version
	c := exec.Command("freebsd-version")
	stdout, err := c.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "While executing external command: %v", err)
		os.Exit(ExitFailure)
	}

	fmt.Print(string(stdout))
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
