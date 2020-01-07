package util_test

import (
	"testing"

	"github.com/tobiashienzsch/jailer/util"
)

func TestCleanString(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"trailing_whitespace", "12.1 ", "12.1 "},
		{"leading_whitespace", " 12.1", "12.1"},
		{"newline", "test\n\r", "test"},
	}

	for _, tt := range tests {

		t.Run(tt.input, func(t *testing.T) {

			actual := util.CleanString(tt.input)
			if actual != tt.expected {
				t.Errorf("Expected: %q, Got: %q", tt.expected, actual)
			}

		})

	}

}
