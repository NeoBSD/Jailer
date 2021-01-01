package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/NeoBSD/jailer"
)

// initCmd represents the config sub command
var initCmd = &cobra.Command{
	Use:   "init [zpool]",
	Short: "Init for jailer. Creates zfs datasets",
	Args:  cobra.MinimumNArgs(1),
	RunE:  RunInitCommand,
}

// RunInitCommand ...
func RunInitCommand(cmd *cobra.Command, args []string) error {

	// Create datasets
	jailerRoot := fmt.Sprintf("%s/jailer", args[0])
	err := createDatasets(jailerRoot)
	if err != nil {
		return err
	}

	// Write initial config
	// Enable jails in rc.conf

	return nil
}
func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

}

func createDatasets(root string) error {
	// Create root dataset
	_, err := jailer.NewZFSDataset(root)
	if err != nil {
		return err
	}
	logrus.Info("Root dataset created")

	// Create config dataset. Holds jail config files
	configRoot := fmt.Sprintf("%s/config", root)
	_, err = jailer.NewZFSDataset(configRoot)
	if err != nil {
		return err
	}
	logrus.Info("Config dataset created")

	// Create base dataset. Holds base FreeBSD image
	baseRoot := fmt.Sprintf("%s/base", root)
	_, err = jailer.NewZFSDataset(baseRoot)
	if err != nil {
		return err
	}
	logrus.Info("Base dataset created")

	// Create jails dataset. Storage for jails
	jailsRoot := fmt.Sprintf("%s/jails", root)
	_, err = jailer.NewZFSDataset(jailsRoot)
	if err != nil {
		return err
	}
	logrus.Info("Jails dataset created")

	return nil
}
