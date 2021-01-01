package jailer

import (
	"strings"
)

// ParseJLSOutput loads the output from `jls` to map
func ParseJLSOutput(source string) (map[string]string, error) {
	// Early return if source is empty
	if len(source) == 0 {
		return nil, nil
	}

	// Create result map
	result := make(map[string]string)

	// Split pairs
	if !strings.ContainsAny(source, " ") {
		key, val := SplitKeyAndValue(source)
		result[key] = val
		return result, nil
	}

	// For each key-value pair
	pairs := strings.Split(source, " ")
	for _, pair := range pairs {
		// Split into key & value
		key, val := SplitKeyAndValue(pair)
		result[key] = val
	}

	return result, nil
}

// SplitKeyAndValue ...
func SplitKeyAndValue(input string) (string, string) {
	elm := strings.Split(input, "=")
	key := elm[0]
	val := elm[1]

	return key, val
}
