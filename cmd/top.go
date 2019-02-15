package cmd

import (
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// topCmd represents the config sub command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Top inside a container",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Start a process:
		c := exec.Command("top", "-jb", "2")
		logrus.Warn(c.Args)
		if err := c.Start(); err != nil {
			logrus.Fatal(err)
		}

		// Wait for the process to finish or kill it after a timeout:
		done := make(chan error, 1)
		go func() {
			done <- c.Wait()
		}()
		select {
		case <-time.After(3 * time.Second):
			if err := c.Process.Kill(); err != nil {
				logrus.Fatal("failed to kill process: ", err)
			}
			logrus.Println("process killed as timeout reached")
		case err := <-done:
			if err != nil {
				logrus.Fatalf("process finished with error = %v", err)
			}
			logrus.Print("process finished successfully")
		}

	},
}

func init() {
	rootCmd.AddCommand(topCmd)

	// Here you will define your flags and configuration settings.

}
