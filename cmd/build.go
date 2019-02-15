package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tobiashienzsch/jailer/jail"
	"github.com/tobiashienzsch/jailer/runtime"
)

// buildCmd represents the config sub command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a jailer container",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		container := jail.Jail{
			Name:      "test-jail",
			Hostname:  "test.jailer",
			IP:        "10.23.0.55",
			Path:      fmt.Sprintf("%s/%s", runtime.JailerPath, "test-jail"),
			ExecStart: "/bin/sh /etc/rc",
			ExecStop:  "/bin/sh /etc/rc.shutdown",
		}
		err := jail.WriteConfig(container)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

}
