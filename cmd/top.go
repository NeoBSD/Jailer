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
	Run: func(cmd *cobra.Command, args []string) {
		externalCMD := "top"
		c := exec.Command(externalCMD, "-jb", "2")

		stdout, err := c.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print(string(stdout))

	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}
