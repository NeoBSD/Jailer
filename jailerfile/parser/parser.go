package parser

import (
	"bufio"
	"os"
	"strings"
)

// Define constants for the command strings
const (
	Add        = "ADD"
	Cmd        = "CMD"
	Copy       = "COPY"
	Entrypoint = "ENTRYPOINT"
	Maintainer = "MAINTAINER"
	From       = "FROM"
	Run        = "RUN"
	Shell      = "SHELL"
)

// Command map of all Jailerfile commands
type Command struct {
	Add        string
	Cmd        string
	Copy       string
	Entrypoint string
	From       string
	Run        string
	Shell      string
	Maintainer string
}

func (c *Command) String() string {
	return "HHIIO"
}

// NewJailerCommandFromFile ...
func NewJailerCommandFromFile(path string) (Command, error) {

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return Command{}, err
	}
	defer file.Close()

	// Scan rows
	command := Command{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split elements by space char
		line := scanner.Text()
		elements := strings.Split(line, " ")

		// Map to config
		switch elements[0] {
		case Add:
			str := strings.Replace(line, elements[0], "", 1)
			command.Add = cleanString(str)
		case Cmd:
			str := strings.Replace(line, elements[0], "", 1)
			command.Cmd = cleanString(str)
		case Copy:
			str := strings.Replace(line, elements[0], "", 1)
			command.Copy = cleanString(str)
		case Entrypoint:
			str := strings.Replace(line, elements[0], "", 1)
			command.Entrypoint = cleanString(str)
		case From:
			str := strings.Replace(line, elements[0], "", 1)
			command.From = cleanString(str)
		case Run:
			str := strings.Replace(line, elements[0], "", 1)
			command.Run = cleanString(str)
		case Shell:
			str := strings.Replace(line, elements[0], "", 1)
			command.Shell = cleanString(str)
		case Maintainer:
			str := strings.Replace(line, elements[0], "", 1)
			command.Maintainer = cleanString(str)
		default:
		}
	}
	// Return empty config & error if an error happend during scanning
	if err := scanner.Err(); err != nil {
		return Command{}, err
	}

	return command, nil
}

func cleanString(str string) string {
	cleaned := strings.Replace(str, " ", "", -1)
	cleaned = strings.TrimLeft(str, " ")
	cleaned = strings.Replace(cleaned, "\n", "", -1)
	cleaned = strings.Replace(cleaned, "\r", "", -1)
	return cleaned
}
