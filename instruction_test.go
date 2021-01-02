package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
	"github.com/matryer/is"
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
			is := is.New(t)
			is.Equal(tt.input.Name(), tt.expected)
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
			is := is.New(t)
			copy := &jailer.CopyInstruction{}
			is.NoErr(copy.Parse(tt.input))
			is.Equal(copy.Command, tt.expected)
		})

	}
}
