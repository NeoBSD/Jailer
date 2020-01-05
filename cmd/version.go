package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tobiashienzsch/jailer/runtime"
)

// versionCmd represents the version sub command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current version",
	Long:  ``,
	Run:   RunVersionCommand,
}

// RunVersionCommand prints the current jailer version
func RunVersionCommand(cmd *cobra.Command, args []string) {
	if viper.Get("verbose") == true {
		fmt.Printf("Version: %s\n", runtime.Version)
		fmt.Printf("Commit: %s\n", runtime.BuildCommit)
		fmt.Printf("Date: %s\n", runtime.BuildDate)
		fmt.Printf("Build on: %s\n", runtime.BuildOS)
		return
	}

	fmt.Printf("%s-%s\n", runtime.Version, runtime.BuildCommit)
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
