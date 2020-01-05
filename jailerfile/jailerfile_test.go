package jailerfile_test

import (
	"testing"

	"github.com/tobiashienzsch/jailer/jailerfile"
)

func TestNewFromFile(t *testing.T) {

	var tests = []struct {
		input       string
		expected    *jailerfile.Jailerfile
		expectError bool
	}{
		// 1
		{"testdata/Jailerfile_noexist", &jailerfile.Jailerfile{}, true},

		// 2
		{"testdata/Jailerfile_basic",
			&jailerfile.Jailerfile{
				BaseImage: jailerfile.BaseImage{
					Name:    "freebsd",
					Version: "latest",
				},
			},
			false,
		},

		// 3
		{"testdata/Jailerfile_multi_string",
			&jailerfile.Jailerfile{
				BaseImage: jailerfile.BaseImage{
					Name:    "freebsd",
					Version: "latest",
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		actual, err := jailerfile.NewFromFile(tt.input)

		// error
		if tt.expectError != (err != nil) {
			t.Errorf("Error %s", err)
		}

		// labels
		// for idx, _ := range actual.Labels {
		// 	if actual.Labels[idx] != tt.expected.Labels[idx] {
		// 		t.Errorf("Expected: %q, Got: %q", tt.expected.Labels[idx], actual.Labels[idx])
		// 	}
		// }

		// from
		if actual.BaseImage != tt.expected.BaseImage {
			t.Errorf("Expected: %q, Got: %q", tt.expected.BaseImage, actual.BaseImage)
		}
	}
}
