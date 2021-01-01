package main

import (
	"testing"
)

func TestRunVersionCommand(t *testing.T) {
	err := RunVersionCommand(nil, []string{})
	if err != nil {
		t.Errorf("Error in version subcommand: %v", err)
	}
}
