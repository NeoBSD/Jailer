package parser_test

import (
	"testing"

	"github.com/tobiashienzsch/jailer/jailerfile/parser"
)

func TestNewCommandFromFile(t *testing.T) {

	var tests = []struct {
		input        string
		expected     parser.Command
		expectsError bool
	}{
		{"testdata/Jailerfile_noexist", parser.Command{}, true},
		{"testdata/Jailerfile_basic", parser.Command{Run: "echo", From: "freebsd"}, false},
		{"testdata/Jailerfile_multi_string", parser.Command{Run: "echo $PWD", From: "freebsd"}, false},
	}

	for _, tt := range tests {
		actual, err := parser.NewCommandFromFile(tt.input)
		if err != nil {
			// If an error exists, but it should not
			if tt.expectsError == false {
				t.Errorf("Error %s", err)
			}
		}
		// Test run command
		if actual.Run != tt.expected.Run {
			t.Errorf("Expected: %q, Got: %q", tt.expected.Run, actual.Run)
		}
		// Test From
		if actual.From != tt.expected.From {
			t.Errorf("Expected: %q, Got: %q", tt.expected.From, actual.From)
		}
	}
}
