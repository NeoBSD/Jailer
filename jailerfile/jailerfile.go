package jailerfile

import (
	"bufio"
	"os"
	"strings"
)

// Jailerfile ...
type Jailerfile struct {
	Maintainer   string
	BaseImage    string
	Instructions []interface{ Instruction }
}

func (j Jailerfile) String() string {
	return j.BaseImage
}

// ParseFromFile parses a Jailerfile from the filesystem
func ParseFromFile(path string) (*Jailerfile, error) {

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return &Jailerfile{}, err
	}

	defer file.Close()

	// Scan rows
	result := &Jailerfile{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split elements by space character
		line := scanner.Text()
		elements := strings.Split(line, " ")
		str := strings.Replace(line, elements[0], "", 1)

		// Switch on instruction keyword
		switch elements[0] {
		case Maintainer:
			result.Maintainer = cleanString(str)
		case From:
			result.BaseImage = cleanString(str)
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
