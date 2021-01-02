package main

import (
	"testing"
)

func TestParseComposeFile(t *testing.T) {
	path := "testdata/jailer-compose.yml"
	compose, err := parseComposeFile(path)
	if err != nil {
		t.Errorf("Error in parseComposeFile subcommand: %v", err)
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
