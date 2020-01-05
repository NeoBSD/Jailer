package jailerfile_test

import (
	"testing"

	"github.com/tobiashienzsch/jailer/jailerfile"
)

func TestInstructionNameToKeywordMapping(t *testing.T) {

	var tests = []struct {
		name     string
		input    jailerfile.Instruction
		expected string
	}{
		{"copy", &jailerfile.CopyInstruction{}, jailerfile.Copy},
		{"cmd", &jailerfile.CmdInstruction{}, jailerfile.Cmd},
		{"from", &jailerfile.FromInstruction{}, jailerfile.From},
		{"run", &jailerfile.RunInstruction{}, jailerfile.Run},
		{"workdir", &jailerfile.WorkDirInstruction{}, jailerfile.WorkDir},
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
			copy := &jailerfile.CopyInstruction{}
			copy.Parse(tt.input)
			actual := copy.Command
			if actual != tt.expected {
				t.Errorf("Expected: %q, Got: %q", tt.expected, actual)
			}
		})

	}
}
