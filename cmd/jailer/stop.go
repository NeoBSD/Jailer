package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// stopCmd represents the `stop` sub command
var stopCmd = &cobra.Command{
	Use:   "stop [jail_id/jail_name]",
	Short: "Stop one or more running jails",
	Long:  ``,
	RunE:  RunStopCommand,
}

// RunStopCommand stoputes the `stop` subcommand.
func RunStopCommand(cmd *cobra.Command, args []string) error {
	id := args[0]
	stop := exec.Command("service", "jail", "stop", id)
	stdout, err := stop.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(stdout))
	return nil
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
