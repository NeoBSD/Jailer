package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the `start` sub command
var startCmd = &cobra.Command{
	Use:   "start JAILS [JAILS...]",
	Short: "Start one or more stopped jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  RunStartCommand,
}

// RunStartCommand startutes the `start` subcommand.
func RunStartCommand(cmd *cobra.Command, args []string) error {
	// Convert arg slice to variadic pack
	cmdArgs := []string{"jail", "start"}
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
	rootCmd.AddCommand(startCmd)
}
