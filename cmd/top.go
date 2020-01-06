package cmd

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// topCmd represents the config sub command
var topCmd = &cobra.Command{
	Use:   "top [jail_id/jail_name]",
	Short: "Run top inside a jail",
	Args:  cobra.MinimumNArgs(1),
	RunE:  RunTopCommand,
}

// RunTopCommand ...
func RunTopCommand(cmd *cobra.Command, args []string) error {

	// -j
	// Display the jail(8) ID.

	// -J jail
	// Show only those processes owned by jail.  This may be either the
	// jid or name of the jail.  Use 0 to limit to host processes.
	// Using this option implies -j.

	// -b
	// Use “batch” mode.  In this mode, all input from the terminal is
	// ignored.  Interrupt characters (such as ^C and ^\) still have an
	// effect.  This is the default on a dumb terminal, or when the
	// output is not a terminal.

	if len(args) < 1 {
		errors.New("requires a jail id/name argument")
	}

	externalCMD := "top"
	c := exec.Command(externalCMD, "-J", args[0], "-b")

	stdout, err := c.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(stdout))

}

func init() {
	rootCmd.AddCommand(topCmd)
}
