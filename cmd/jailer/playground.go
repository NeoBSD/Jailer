package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NeoBSD/jailer"
	"github.com/spf13/cobra"
)

// playgroundCmd represents the `playground` sub command
var playgroundCmd = &cobra.Command{
	Use:   "playground",
	Short: "Subcommand for development only",
	Run:   RunPlaygroundCommand,
}

// RunPlaygroundCommand executes the playground subcommand.
func RunPlaygroundCommand(cmd *cobra.Command, args []string) {
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
	rootCmd.AddCommand(playgroundCmd)
}
