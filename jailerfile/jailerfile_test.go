package jailerfile_test

import (
	"testing"

	"github.com/tobiashienzsch/jailer/jailerfile"
)

func TestParseFromFile(t *testing.T) {

	var tests = []struct {
		input       string
		expected    *jailerfile.Jailerfile
		expectError bool
	}{
		{"testdata/Jailerfile_noexist", &jailerfile.Jailerfile{}, true},
		{"testdata/Jailerfile_basic", &jailerfile.Jailerfile{BaseImage: "freebsd"}, false},
		{"testdata/Jailerfile_multi_string", &jailerfile.Jailerfile{BaseImage: "freebsd"}, false},
	}

	for _, tt := range tests {
		actual, err := jailerfile.ParseFromFile(tt.input)

		// error
		if tt.expectError != (err != nil) {
			t.Errorf("Error %s", err)
		}

		// maintainer
		if actual.Maintainer != tt.expected.Maintainer {
			t.Errorf("Expected: %q, Got: %q", tt.expected.Maintainer, actual.Maintainer)
		}

		// from
		if actual.BaseImage != tt.expected.BaseImage {
			t.Errorf("Expected: %q, Got: %q", tt.expected.BaseImage, actual.BaseImage)
		}
	}
}
