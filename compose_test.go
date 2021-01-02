package jailer_test

import (
	"testing"

	"github.com/NeoBSD/jailer"
)

func TestReadComposeFileFail(t *testing.T) {
	_, err := jailer.ReadComposeFile("unknown")
	if err == nil {
		t.Errorf("expected error wrong filepath, got no error")
	}
}

func TestReadComposeFile(t *testing.T) {
	path := "testdata/jailer-compose/jailer-compose.yml"
	compose, err := jailer.ReadComposeFile(path)
	if err != nil {
		t.Errorf("Error in ReadComposeFile subcommand: %v", err)
	}

	if compose.Version != "0.1" {
		t.Errorf("Expected: version %s, got %s", "0.1", compose.Version)
	}

	if len(compose.Services) != 2 {
		t.Errorf("Expected: %d services, got %d", 2, len(compose.Services))
	}

	if compose.Services[0].Label != "web" {
		t.Errorf("Expected: version %s, got %s", "web", compose.Services[0].Label)
	}

}

func TestComposeValidate(t *testing.T) {
	// OK
	{
		c := jailer.Compose{
			Version: "0.1",
			Services: []jailer.Service{
				{Label: "test1"},
				{Label: "test2"},
			},
		}
		if err := c.Validate(); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	}

	// FAIL: Missing version
	{
		c := jailer.Compose{}
		if err := c.Validate(); err == nil {
			t.Errorf("expected validation fail, because of missing version")
		}
	}

	// FAIL: Label used twice
	{
		c := jailer.Compose{
			Version: "0.1",
			Services: []jailer.Service{
				{Label: "test1"},
				{Label: "test1"},
			},
		}
		if err := c.Validate(); err == nil {
			t.Errorf("expected validation fail, label used twice")
		}
	}
}
