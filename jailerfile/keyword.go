package jailerfile

import (
	"strings"
)

// Define constants for the command strings
const (
	// Add        = "ADD"
	// Entrypoint = "ENTRYPOINT"
	Cmd     = "CMD"
	Copy    = "COPY"
	From    = "FROM"
	Label   = "LABEL"
	Run     = "RUN"
	Shell   = "SHELL"
	WorkDir = "WORKDIR"
)

func cleanString(input string) string {
	cleaned := strings.Replace(input, " ", "", -1)
	cleaned = strings.TrimLeft(input, " ")
	cleaned = strings.Replace(cleaned, "\n", "", -1)
	cleaned = strings.Replace(cleaned, "\r", "", -1)
	return cleaned
}
