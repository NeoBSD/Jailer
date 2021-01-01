package main

import (
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// storageCmd represents the storage sub command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manages jail & image storage",
	Long:  ``,
	Run:   RunStorageCommand,
}

// RunStorageCommand ...
func RunStorageCommand(cmd *cobra.Command, args []string) {

	// -H
	// Used for scripting mode. Do not print headers and separate
	// fields by a single tab instead of arbitrary white space.

	// -o all | field[,field]...
	// A comma-separated list of columns to display. Supported
	// values are name,property,value,received,source.  Default
	// values are name,property,value,source.  The keyword all
	// specifies all columns.

	// Setup external zfs list
	externalCMD := "zfs"
	c := exec.Command(externalCMD, "list", "-H", "-o", "name")

	// Exec external zfs list
	stdout, err := c.Output()
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.Print(string(stdout))

}

func init() {
	rootCmd.AddCommand(storageCmd)

	// Here you will define your flags and configuration settings.

}
