package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/NeoBSD/jailer"
	"github.com/NeoBSD/jailer/freebsd"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	if viper.GetBool("json") {
		js, err := json.Marshal(host)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(js))
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, tabWriterPadding, ' ', 0)

	v := host.Version
	fmt.Fprintf(w, "Hostname\t%s\t\n", host.Hostname)
	fmt.Fprintf(w, "OS\t%d.%d-%s-p%d\t\n", v.Major, v.Minor, v.Comment, v.Patch)

	fmt.Fprintf(w, "Version\t%s\t\n", jailer.Version)
	fmt.Fprintf(w, "Date\t%s\t\n", jailer.BuildDate)
	fmt.Fprintf(w, "Commit\t%s\t\n", jailer.BuildCommit)

	fmt.Fprintf(w, "Machine\t%s\t\n", host.Machine)
	fmt.Fprintf(w, "MachineArch\t%s\t\n", host.MachineArch)
	fmt.Fprintf(w, "Model\t%s\t\n", host.Hardware.Model)
	fmt.Fprintf(w, "NumCPU\t%d\t\n", host.Hardware.NumCPU)

	w.Flush()
	return nil
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
