package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// topCmd represents the config sub command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Top inside a container",
	Long:  ``,
	Run:   RunTopCommand,
}

// RunTopCommand ...
func RunTopCommand(cmd *cobra.Command, args []string) {

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

	externalCMD := "top"
	c := exec.Command(externalCMD, "-jb", "2")

	stdout, err := c.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))

}

func init() {
	rootCmd.AddCommand(topCmd)
}
