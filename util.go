package jailer

import "strings"

// CleanString removes leading spaces & newlines in a string
func CleanString(input string) string {
	cleaned := strings.TrimLeft(input, " ")
	cleaned = strings.Replace(cleaned, "\n", "", -1)
	cleaned = strings.Replace(cleaned, "\r", "", -1)
	return cleaned
}
