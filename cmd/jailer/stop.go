package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// stopCmd represents the `stop` sub command
var stopCmd = &cobra.Command{
	Use:   "stop JAILS [JAILS...]",
	Short: "Stop one or more running jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  RunStopCommand,
}

// RunStopCommand stoputes the `stop` subcommand.
func RunStopCommand(cmd *cobra.Command, args []string) error {
	// Convert arg slice to variadic pack
	cmdArgs := []string{"jail", "stop"}
	cmdArgs = append(cmdArgs, args...)

	restart := exec.Command("service", cmdArgs...)
	stdout, err := restart.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(stdout))
	return nil
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
