package cmd

import (
	"os"
	"os/user"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:              "jailer",
	Short:            "jailer",
	Long:             `jailer https://github.com/NeoBSD/jailer`,
	PersistentPreRun: PreRunRootCommand,
	Run:              RunRootCommand,
}

// PreRunRootCommand gets called before every command & subcommand
func PreRunRootCommand(cmd *cobra.Command, args []string) {
	// Skip root user check for 'version' subcommand
	if cmd.Name() == "version" || cmd.Name() == "playground" {
		return
	}

	// Get current os user
	user, err := user.Current()
	if err != nil {
		logrus.Warning("Could not get current user. Jailer should run as root")
	}

	// Check if jailer is running as root. Abort if not.
	if user.Username != "root" {
		logrus.Fatalf("Jailer should run as root. You are %v. Switch user or use sudo.", user.Username)
	}
}

// RunRootCommand runs if no subcommand was selected.
func RunRootCommand(cmd *cobra.Command, args []string) {
}

// Execute is the main entry point for the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}

func init() {
	// Viper config
	cobra.OnInitialize(initConfig)

	// Flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $PWD/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find pwd
		dir, err := os.Getwd()
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("config")
		viper.AutomaticEnv() // read in environment variables that match

	}

	// Set log level
	// TODO: set via config.yml
	if viper.Get("verbose") == true {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Can't read config: %v\n", err)
		os.Exit(1)
	}
}
