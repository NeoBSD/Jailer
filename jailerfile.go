package jailer

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
	Image        string            `json:"image"`
	Labels       map[string]string `json:"labels"`
	BaseImage    BaseImage         `json:"base_image"`
	Instructions []Instruction     `json:"instructions"`
}

func (j Jailerfile) String() string {
	return fmt.Sprintf("%s:%s", j.BaseImage.Name, j.BaseImage.Version)
}

// ReadFromFile parses a Jailerfile from the filesystem
func ReadFromFile(path string) (*Jailerfile, error) {

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
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
			key := CleanString(label[0])
			value := CleanString(label[1])
			result.Labels[key] = value
		case From:
			base := strings.Split(str, ":")
			name := CleanString(base[0])
			result.BaseImage.Name = name
			result.BaseImage.Version = "latest"

			if len(base) == 2 {
				result.BaseImage.Version = CleanString(base[1])
			}
			from := &FromInstruction{From: fmt.Sprintf("%s:%s", name, result.BaseImage.Version)}
			result.Instructions = append(result.Instructions, from)
		case Run:
			run := &RunInstruction{}
			err = parseInstructionLine(result, run, str)
		case WorkDir:
			workDir := &WorkDirInstruction{}
			err = parseInstructionLine(result, workDir, str)
		case Copy:
			copy := &CopyInstruction{}
			err = parseInstructionLine(result, copy, str)
		case Cmd:
			cmd := &CmdInstruction{}
			err = parseInstructionLine(result, cmd, str)
		default:
		}
	}

	if err != nil {
		return nil, err
	}

	// Return empty & error if an error happend during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func parseInstructionLine(jf *Jailerfile, data interface{ Instruction }, line string) error {

	err := data.Parse(line)
	if err != nil {
		return err
	}

	jf.Instructions = append(jf.Instructions, data)
	return nil
}
