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
		{"testdata/noexist/Jailerfile", &jailerfile.Jailerfile{}, true},
		{"testdata/label/Jailerfile", &jailerfile.Jailerfile{BaseImage: jailerfile.BaseImage{Name: "freebsd", Version: "latest"}}, false},
	}

	for _, tt := range tests {
		actual, err := jailerfile.NewFromFile(tt.input)

		// error
		if tt.expectError != (err != nil) {
			t.Errorf("Error %s", err)
		}

		// from
		if actual.BaseImage != tt.expected.BaseImage {
			t.Errorf("Expected: %q, Got: %q", tt.expected.BaseImage, actual.BaseImage)
		}
	}
}

func TestLabelParsing(t *testing.T) {

	jf, err := jailerfile.NewFromFile("testdata/label/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.Labels["maintainer"] != `"example@example.com"` {
		t.Errorf("Expected: \"%s\", got %s", "example@example.com", jf.Labels["maintainer"])
	}

	if jf.Labels["version"] != `"1.0"` {
		t.Errorf("Expected: \"%s\", got %s", "1.0", jf.Labels["version"])
	}

}

func TestFromWithImplicitLatestParsing(t *testing.T) {

	jf, err := jailerfile.NewFromFile("testdata/from/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.BaseImage.Name != "freebsd" {
		t.Errorf("Expected: %s, got %s", "freebsd", jf.BaseImage.Name)
	}

	if jf.BaseImage.Version != "latest" {
		t.Errorf("Expected: %s, got %s", "latest", jf.BaseImage.Version)
	}

}

func TestFromWithExplicitLatestParsing(t *testing.T) {

	jf, err := jailerfile.NewFromFile("testdata/from_with_latest/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.BaseImage.Name != "freebsd" {
		t.Errorf("Expected: %s, got %s", "freebsd", jf.BaseImage.Name)
	}

	if jf.BaseImage.Version != "latest" {
		t.Errorf("Expected: %s, got %s", "latest", jf.BaseImage.Version)
	}

}

func TestFromWithExplicitVersionParsing(t *testing.T) {

	jf, err := jailerfile.NewFromFile("testdata/from_with_version/Jailerfile")

	if err != nil {
		t.Errorf("Error %v", err)
	}

	if jf.BaseImage.Name != "freebsd" {
		t.Errorf("Expected: %s, got %s", "freebsd", jf.BaseImage.Name)
	}

	if jf.BaseImage.Version != "12.1" {
		t.Errorf("Expected: %s, got %s", "12.1", jf.BaseImage.Version)
	}

}
