package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// execCmd represents the `exec` sub command
var execCmd = &cobra.Command{
	Use:   "exec [jail_id/jail_name] command",
	Short: "Execute a command inside an existing jail",
	Long:  ``,
	RunE:  RunExecCommand,
}

// RunExecCommand executes the `exec` subcommand.
func RunExecCommand(cmd *cobra.Command, args []string) error {
	id := args[0]
	command := args[1]
	jexec := exec.Command("jexec", "-l", id, command)
	stdout, err := jexec.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(stdout))
	return nil
}

func init() {
	rootCmd.AddCommand(execCmd)
}
