package jailerfile

import (
	"strings"
)

// Define constants for the command strings
const (
	Add        = "ADD"
	Cmd        = "CMD"
	Copy       = "COPY"
	Entrypoint = "ENTRYPOINT"
	Label      = "LABEL"
	From       = "FROM"
	Run        = "RUN"
	Shell      = "SHELL"
)

func cleanString(input string) string {
	cleaned := strings.Replace(input, " ", "", -1)
	cleaned = strings.TrimLeft(input, " ")
	cleaned = strings.Replace(cleaned, "\n", "", -1)
	cleaned = strings.Replace(cleaned, "\r", "", -1)
	return cleaned
}
