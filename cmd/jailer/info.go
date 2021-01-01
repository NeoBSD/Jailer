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
	fbsd, err := freebsd.GetSystemInfo()
	if err != nil {
		return err
	}

	// Jailer
	fmt.Println("Jailer:")
	fmt.Printf("\tVersion: %s\n", jailer.Version)
	fmt.Printf("\tDate: %s\n", jailer.BuildDate)
	fmt.Printf("\tCommit: %s\n", jailer.BuildCommit)

	// Hardware
	fmt.Println("Hardware:")
	fmt.Printf("\tMachine: %s\n", fbsd.Machine)
	fmt.Printf("\tMachineArch: %s\n", fbsd.MachineArch)
	fmt.Printf("\tModel: %s\n", fbsd.Hardware.Model)
	fmt.Printf("\tNumCPU: %d\n", fbsd.Hardware.NumCPU)

	// Operating System
	v := fbsd.Version
	fmt.Println("Operating System:")
	fmt.Printf("\tVersion: %d.%d-%s-p%d\n", v.Major, v.Minor, v.Comment, v.Patch)
	return nil
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
