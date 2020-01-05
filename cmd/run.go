package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tobiashienzsch/jailer/jailerfile"
)

// runCmd represents the config sub command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a jailer container",
	Long:  ``,
	Run:   RunRunCommand,
}

// RunRunCommand ...
func RunRunCommand(cmd *cobra.Command, args []string) {

	// Find home directory.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jailFile, err := jailerfile.ParseFromFile(filepath.Join(dir, "testdata", "Jailerfile"))
	if err != nil {
		logrus.Fatal(err)
	}

	// Start a process:
	externalCMD := exec.Command("cat", jailFile.String())
	logrus.Warn(externalCMD.Args)
	if err := externalCMD.Start(); err != nil {
		logrus.Fatal(err)
	}

	// Wait for the process to finish or kill it after a timeout:
	done := make(chan error, 1)
	go func() {
		done <- externalCMD.Wait()
	}()
	select {
	case <-time.After(3 * time.Second):
		if err := externalCMD.Process.Kill(); err != nil {
			logrus.Fatal("failed to kill process: ", err)
		}
		logrus.Println("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			logrus.Fatalf("process finished with error = %v", err)
		}
		logrus.Print("process finished successfully")
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

}
