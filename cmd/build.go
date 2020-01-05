package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tobiashienzsch/jailer/jail"
)

// buildCmd represents the `build` sub command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a jailer container",
	Long:  ``,
	Run:   RunBuildCommand,
}

// RunBuildCommand executes the build subcommand.
func RunBuildCommand(cmd *cobra.Command, args []string) {
	jailerPath := viper.Get("jailer-path")
	name := "test-jail"
	ip := "10.23.0.55"

	// Create config
	container := jail.Jail{
		Name:      name,
		Hostname:  fmt.Sprintf("%s.jailer.com", name),
		IP:        ip,
		Path:      fmt.Sprintf("%s/%s", jailerPath, name),
		ExecStart: "/bin/sh /etc/rc",
		ExecStop:  "/bin/sh /etc/rc.shutdown",
	}

	// Write to *.conf file
	err := jail.WriteConfig(container)
	if err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
