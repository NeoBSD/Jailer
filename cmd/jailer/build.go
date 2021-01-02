package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NeoBSD/jailer"
	"github.com/spf13/cobra"
)

// buildCmd represents the `build` sub command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build an image from a Jailerfile",
	Args:  cobra.ExactArgs(1),
	RunE:  RunBuildCommand,
}

// RunBuildCommand executes the build subcommand.
func RunBuildCommand(cmd *cobra.Command, args []string) error {
	path := fmt.Sprintf("%s/Jailerfile", args[0])
	jf, err := jailer.ReadFromFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	js, err := json.Marshal(jf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(js))

	return nil
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
