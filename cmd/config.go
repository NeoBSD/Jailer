package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config sub command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print current config",
	Long:  ``,
	Run:   RunConfigCommand,
}

// RunConfigCommand ...
func RunConfigCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("Verbose: %v\n", viper.Get("verbose"))

}
func init() {
	rootCmd.AddCommand(configCmd)
}
