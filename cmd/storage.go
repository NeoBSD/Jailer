package cmd

import (
	"fmt"
	"os/exec"
	_ "strings"

	"github.com/spf13/cobra"
)

// storageCmd represents the config sub command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manages container & image storage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Setup external zfs list
		externalCMD := "zfs"
		c := exec.Command(externalCMD, "list", "-H", "-o", "name")

		// Exec external zfs list
		stdout, err := c.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print(string(stdout))

	},
}

func init() {
	rootCmd.AddCommand(storageCmd)

	// Here you will define your flags and configuration settings.

}
