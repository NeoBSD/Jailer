package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tobiashienzsch/jailer/filesystem"
)

// initCmd represents the config sub command
var initCmd = &cobra.Command{
	Use:   "init [zpool]",
	Short: "Init for jailer. Creates zfs datasets",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		// Create datasets
		jailerRoot := fmt.Sprintf("%s/jailer", args[0])
		createDatasets(jailerRoot)

		// Write initial config
		// Enable jails in rc.conf

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

}

func createDatasets(root string) {
	// Create root dataset
	err := filesystem.NewDataset(root)
	if err != nil {
		logrus.Fatalf("Error while creating the root dataset: %v", err)
	}
	logrus.Info("Root dataset created")

	// Create config dataset. Holds jail config files
	configRoot := fmt.Sprintf("%s/config", root)
	err = filesystem.NewDataset(configRoot)
	if err != nil {
		logrus.Fatalf("Error while creating the config dataset: %v", err)
	}
	logrus.Info("Config dataset created")

	// Create base dataset. Holds base FreeBSD image
	baseRoot := fmt.Sprintf("%s/base", root)
	err = filesystem.NewDataset(baseRoot)
	if err != nil {
		logrus.Fatalf("Error while creating the base dataset: %v", err)
	}
	logrus.Info("Base dataset created")

	// Create containers dataset. Storage for containers
	containersRoot := fmt.Sprintf("%s/containers", root)
	err = filesystem.NewDataset(containersRoot)
	if err != nil {
		logrus.Fatalf("Error while creating the config dataset: %v", err)
	}
	logrus.Info("Containers dataset created")
}
