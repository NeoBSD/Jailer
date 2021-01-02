package main

import (
	"fmt"

	"github.com/NeoBSD/jailer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// validateCmd represents the `validate` sub command
var validateCmd = &cobra.Command{
	Use:   "validate path/to/jailer-compose.yml",
	Short: "Validate a jailer-compose.yml file",
	Args:  cobra.ExactArgs(0),
	RunE:  RunValidateCommand,
}

// RunValidateCommand executes the `validate` subcommand.
func RunValidateCommand(cmd *cobra.Command, args []string) error {
	path := viper.GetString("file")
	compose, err := jailer.ReadComposeFile(path)
	if err != nil {
		return err
	}

	validationErrors := compose.Validate()
	if validationErrors != nil {
		return err
	}

	if viper.GetBool("verbose") {
		fmt.Println("ok")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
