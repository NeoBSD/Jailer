package jailerfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// BaseImage ...
type BaseImage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Jailerfile ...
type Jailerfile struct {
	Labels       map[string]string          `json:"labels"`
	BaseImage    BaseImage                  `json:"base_image"`
	Instructions []interface{ Instruction } `json:"instructions"`
}

func (j Jailerfile) String() string {
	return fmt.Sprintf("%s:%s", j.BaseImage.Name, j.BaseImage.Version)
}

// ParseFromFile parses a Jailerfile from the filesystem
func ParseFromFile(path string) (*Jailerfile, error) {

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return &Jailerfile{}, err
	}

	defer file.Close()

	// Init result
	result := &Jailerfile{}
	result.Labels = make(map[string]string)

	// Scan rows
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split elements by space character
		line := scanner.Text()
		elements := strings.Split(line, " ")
		str := strings.Replace(line, elements[0], "", 1)

		// Switch on instruction keyword
		switch elements[0] {
		case Label:
			label := strings.Split(str, "=")
			key := cleanString(label[0])
			value := cleanString(label[1])
			result.Labels[key] = value
		case From:
			base := strings.Split(str, ":")
			name := cleanString(base[0])
			result.BaseImage.Name = name
			result.BaseImage.Version = "latest"

			if len(base) == 2 {
				result.BaseImage.Version = cleanString(base[1])
			}
		case Run:
			result.Instructions = append(result.Instructions, RunInstruction{Command: cleanString(str)})
		default:
		}
	}

	// Return empty & error if an error happend during scanning
	if err := scanner.Err(); err != nil {
		return &Jailerfile{}, err
	}

	return result, nil
}
