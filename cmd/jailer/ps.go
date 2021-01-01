package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/NeoBSD/jailer"
	"github.com/spf13/cobra"
)

// psCmd represents the `ps` sub command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List active jails",
	RunE:  RunPsCommand,
}

// RunPsCommand executes the `ps` subcommand.
func RunPsCommand(cmd *cobra.Command, args []string) error {
	jls := jailer.JLS{Path: "jls"}
	jails, err := jls.GetActiveJails()
	if err != nil {
		return err
	}

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "JID\tName\tHostname\tPath\t")
	for _, jail := range jails {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t\n", jail.JID, jail.Name, jail.Hostname, jail.Path)
	}
	w.Flush()

	return nil
}

func init() {
	rootCmd.AddCommand(psCmd)
}
