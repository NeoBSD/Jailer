package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config sub command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print current config",
	RunE:  RunConfigCommand,
}

// RunConfigCommand ...
func RunConfigCommand(cmd *cobra.Command, args []string) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, tabWriterPadding, ' ', 0)
	for _, key := range viper.AllKeys() {
		value := viper.Get(key)
		fmt.Fprintf(w, "%s\t%v\t\n", key, value)
	}

	return w.Flush()
}

func init() {
	rootCmd.AddCommand(configCmd)
}
