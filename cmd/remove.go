package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the config sub command
var removeCmd = &cobra.Command{
	Use:   "rm [ID]",
	Short: "Remove a container",
	Args:  cobra.MinimumNArgs(1),
	RunE:  RunRemoveCommand,
}

// RunRemoveCommand ...
func RunRemoveCommand(cmd *cobra.Command, args []string) error {

	// Setup external zfs list
	jailPath := fmt.Sprintf("%s/%s", viper.Get("jailer-path"), args[0])
	externalCMD := "zfs"
	c := exec.Command(externalCMD, "list", "-H", "-o", "name", jailPath)

	// Exec external zfs list
	stdout, err := c.Output()
	if err != nil {
		return fmt.Errorf("Error while executing external command: %v", err.Error())
	}

	fmt.Print(string(stdout))

	return nil
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

}
