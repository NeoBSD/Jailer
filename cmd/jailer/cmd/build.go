package cmd

import (
	"github.com/spf13/cobra"
)

// buildCmd represents the `build` sub command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a jailer jail",
	Long:  ``,
	Run:   RunBuildCommand,
}

// RunBuildCommand executes the build subcommand.
func RunBuildCommand(cmd *cobra.Command, args []string) {
	// jailerPath := viper.Get("jailer-path")
	// name := "test-jail"
	// ip := "10.23.0.55"

	// // Create config
	// j := jailer.Jail{
	// 	Name:      name,
	// 	Hostname:  fmt.Sprintf("%s.jailer.com", name),
	// 	IP:        ip,
	// 	Path:      fmt.Sprintf("%s/%s", jailerPath, name),
	// 	ExecStart: "/bin/sh /etc/rc",
	// 	ExecStop:  "/bin/sh /etc/rc.shutdown",
	// }

	// // Write to *.conf file
	// err := jailer.WriteConfig(j)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%v", err)
	// 	os.Exit(ExitFailure)
	// }
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
