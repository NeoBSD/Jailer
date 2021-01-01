package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the `start` sub command
var startCmd = &cobra.Command{
	Use:   "start [jail_id/jail_name]",
	Short: "Start one or more stopped jails",
	Long:  ``,
	RunE:  RunStartCommand,
}

// RunStartCommand startutes the `start` subcommand.
func RunStartCommand(cmd *cobra.Command, args []string) error {
	id := args[0]
	start := exec.Command("service", "jail", "start", id)
	stdout, err := start.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(stdout))
	return nil
}

func init() {
	rootCmd.AddCommand(startCmd)
}
