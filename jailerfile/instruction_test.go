package jailerfile_test

import (
	"testing"

	"github.com/tobiashienzsch/jailer/jailerfile"
)

func TestFromInstructionName(t *testing.T) {

	var tests = []struct {
		input    jailerfile.Instruction
		expected string
	}{
		{&jailerfile.FromInstruction{}, "FROM"},
		{&jailerfile.RunInstruction{}, "RUN"},
	}

	for _, tt := range tests {
		actual := tt.input.Name()
		if actual != tt.expected {
			t.Errorf("Expected: %q, Got: %q", tt.expected, actual)
		}
	}
}
