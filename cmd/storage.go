package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// storageCmd represents the config sub command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manages container & image storage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		cArgs := "list"
		c := exec.Command("zfs", strings.Split(cArgs, " ")...)

		stderr, _ := c.StdoutPipe()
		c.Start()

		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}
		c.Wait()

	},
}

func init() {
	rootCmd.AddCommand(storageCmd)

	// Here you will define your flags and configuration settings.

}
