package main

import (
	"fmt"

	"github.com/NeoBSD/jailer"
	"github.com/NeoBSD/jailer/freebsd"

	"github.com/spf13/cobra"
)

// infoCmd represents the "info" sub command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display system-wide information",
	RunE:  RunInfoCommand,
}

// RunInfoCommand infoutes the "info" subcommand.
func RunInfoCommand(cmd *cobra.Command, args []string) error {
	host, err := freebsd.GetSystemInfo()
	if err != nil {
		return err
	}

	// Host
	v := host.Version
	fmt.Println("Host:")
	fmt.Printf("\tHostname: %s\n", host.Hostname)
	fmt.Printf("\tOS: %d.%d-%s-p%d\n", v.Major, v.Minor, v.Comment, v.Patch)

	// Jailer
	fmt.Println("Jailer:")
	fmt.Printf("\tVersion: %s\n", jailer.Version)
	fmt.Printf("\tDate: %s\n", jailer.BuildDate)
	fmt.Printf("\tCommit: %s\n", jailer.BuildCommit)

	// Hardware
	fmt.Println("Hardware:")
	fmt.Printf("\tMachine: %s\n", host.Machine)
	fmt.Printf("\tMachineArch: %s\n", host.MachineArch)
	fmt.Printf("\tModel: %s\n", host.Hardware.Model)
	fmt.Printf("\tNumCPU: %d\n", host.Hardware.NumCPU)

	return nil
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
