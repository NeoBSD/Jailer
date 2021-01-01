package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// restartCmd represents the `restart` sub command
var restartCmd = &cobra.Command{
	Use:   "restart JAILS [JAILS...]",
	Short: "Restart one or more jails",
	Args:  cobra.MinimumNArgs(1),
	RunE:  RunRestartCommand,
}

// RunRestartCommand restartutes the `restart` subcommand.
func RunRestartCommand(cmd *cobra.Command, args []string) error {
	// Convert arg slice to variadic pack
	cmdArgs := []string{"jail", "restart"}
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
	rootCmd.AddCommand(restartCmd)
}
