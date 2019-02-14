// Package command contains the list of all possible Jailerfile commands.
package command

// Constant mapping
const (
	Add        = "add"
	Cmd        = "cmd"
	Copy       = "copy"
	Entrypoint = "entrypoint"
	From       = "from"
	Run        = "run"
	Shell      = "shell"
)

// Commands slice of all Jailerfile commands
var Commands = map[string]struct{}{
	Add:        {},
	Cmd:        {},
	Copy:       {},
	Entrypoint: {},
	From:       {},
	Run:        {},
	Shell:      {},
}
