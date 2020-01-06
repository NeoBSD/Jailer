package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tobiashienzsch/jailer/jailerfile"
)

// playgroundCmd represents the `build` sub command
var playgroundCmd = &cobra.Command{
	Use:   "playground",
	Short: "Subcommand for development only",
	Long:  ``,
	Run:   RunPlaygroundCommand,
}

// RunPlaygroundCommand executes the build subcommand.
func RunPlaygroundCommand(cmd *cobra.Command, args []string) {
	jf, err := jailerfile.NewFromFile("example/Jailerfile")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	data, err := json.Marshal(jf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	logrus.Print(string(data))
}

func init() {
	rootCmd.AddCommand(playgroundCmd)
}
