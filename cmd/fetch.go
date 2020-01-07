package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/tobiashienzsch/jailer/jail"
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

	ftpURL := "https://download.freebsd.org/ftp/releases/amd64/12.0-RELEASE/"

	baseName := "base.txz"
	baseURL := fmt.Sprintf("%s%s", ftpURL, baseName)

	if err := jail.DownloadFile(baseName, baseURL); err != nil {
		fmt.Fprintf(os.Stderr, "While downloading: %v", err)
		os.Exit(ExitFailure)
	}
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
