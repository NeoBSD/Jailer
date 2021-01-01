package cmd_test

import (
	"testing"

	"github.com/NeoBSD/jailer/cmd/jailer/cmd"
)

func TestRunVersionCommand(t *testing.T) {
	err := cmd.RunVersionCommand(nil, []string{})
	if err != nil {
		t.Errorf("Error in version subcommand: %v", err)
	}
}
