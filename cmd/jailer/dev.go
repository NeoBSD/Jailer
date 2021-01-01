package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NeoBSD/jailer"
	"github.com/spf13/cobra"
)

// devCmd represents the `dev` sub command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Subcommand for development only",
	Run:   RunDevCommand,
}

// RunDevCommand executes the dev subcommand.
func RunDevCommand(cmd *cobra.Command, args []string) {
	jf, err := jailer.NewFromFile("example/Jailerfile")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	data, err := json.Marshal(jf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(data))

}

func init() {
	rootCmd.AddCommand(devCmd)
}
