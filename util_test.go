package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
	"github.com/matryer/is"
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
		is := is.New(t)
		actual := jailer.CleanString(tt.input)
		is.Equal(actual, tt.expected)
	}

}
