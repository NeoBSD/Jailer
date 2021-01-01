package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
)

func TestInstructionNameToKeywordMapping(t *testing.T) {

	var tests = []struct {
		name     string
		input    jailer.Instruction
		expected string
	}{
		{"copy", &jailer.CopyInstruction{}, jailer.Copy},
		{"cmd", &jailer.CmdInstruction{}, jailer.Cmd},
		{"from", &jailer.FromInstruction{}, jailer.From},
		{"run", &jailer.RunInstruction{}, jailer.Run},
		{"workdir", &jailer.WorkDirInstruction{}, jailer.WorkDir},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.Name()
			if actual != tt.expected {
				t.Errorf("Expected: %q, Got: %q", tt.expected, actual)
			}
		})

	}
}

func TestCopyInstructionParsing(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", ""},
		{"file", " test.txt", "test.txt"},
		{"cwd", " test.txt .", "test.txt ."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copy := &jailer.CopyInstruction{}
			copy.Parse(tt.input)
			actual := copy.Command
			if actual != tt.expected {
				t.Errorf("Expected: %q, Got: %q", tt.expected, actual)
			}
		})

	}
}
