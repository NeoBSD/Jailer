package jailer

import (
	"fmt"
	"strings"
)

// Instruction represents any Jailerfile instruction.
type Instruction interface {
	// Execute() error
	Name() string
	Parse(source string) error
	Execute(jf *Jailerfile) error
}

// FromInstruction specifies te base jail
type FromInstruction struct {
	From string `json:"from"`
}

// String returns the instruction command
func (f FromInstruction) String() string {
	return f.Name()
}

// Name returns the instruction identifier
func (f FromInstruction) Name() string {
	return "FROM"
}

// Parse parses the source string
func (f FromInstruction) Parse(input string) error {
	return nil
}

// Execute runs the instruction on the current Jailerfile
func (f FromInstruction) Execute(jf *Jailerfile) error {
	fmt.Printf("# from=%s\n", jf.BaseImage.Name)
	return nil
}

// RunInstruction executes a command at build time inside the jail
type RunInstruction struct {
	Command    string `json:"run"`
	jailerfile *Jailerfile
}

// String returns the instruction command
func (r RunInstruction) String() string {
	return r.Name()
}

// Name returns the instruction identifier
func (r RunInstruction) Name() string {
	return "RUN"
}

// Parse parses the source string
func (r *RunInstruction) Parse(input string) error {
	r.Command = CleanString(input)
	return nil
}

// Execute runs the instruction on the current Jailerfile
func (r RunInstruction) Execute(jf *Jailerfile) error {
	fmt.Printf("jexec -l %s /bin/sh -c %s\n", jf.Image, r.Command)
	return nil
}

// WorkDirInstruction sets the working directory at build time inside the jail
type WorkDirInstruction struct {
	Command    string `json:"work_dir"`
	jailerfile *Jailerfile
}

// String returns the instruction command
func (w WorkDirInstruction) String() string {
	return w.Name()
}

// Name returns the instruction identifier
func (w WorkDirInstruction) Name() string {
	return "WORKDIR"
}

// Parse parses the source string
func (w *WorkDirInstruction) Parse(input string) error {
	w.Command = CleanString(input)
	return nil
}

// Execute runs the instruction on the current Jailerfile
func (w WorkDirInstruction) Execute(jf *Jailerfile) error {
	fmt.Printf("# workdir=%s\n", w.Command)
	return nil
}

// CopyInstruction sets the working directory at build time inside the jail
type CopyInstruction struct {
	Command    string `json:"copy"`
	jailerfile *Jailerfile
}

// String returns the instruction command
func (c CopyInstruction) String() string {
	return c.Name()
}

// Name returns the instruction identifier
func (c CopyInstruction) Name() string {
	return "COPY"
}

// Parse parses the source string
func (c *CopyInstruction) Parse(input string) error {
	c.Command = CleanString(input)
	return nil
}

// Execute runs the instruction on the current Jailerfile
func (c CopyInstruction) Execute(jf *Jailerfile) error {
	splits := strings.Split(c.Command, " ")
	src := splits[0]
	dest := splits[1]
	fmt.Printf("cp %s %s\n", src, dest)
	return nil
}

// CmdInstruction sets the working directory at build time inside the jail
type CmdInstruction struct {
	Command    string `json:"cmd"`
	jailerfile *Jailerfile
}

// String returns the instruction command
func (c CmdInstruction) String() string {
	return c.Name()
}

// Name returns the instruction identifier
func (c CmdInstruction) Name() string {
	return "CMD"
}

// Parse parses the source string
func (c *CmdInstruction) Parse(input string) error {
	c.Command = CleanString(input)
	return nil
}

// Execute runs the instruction on the current Jailerfile
func (c CmdInstruction) Execute(jf *Jailerfile) error {
	fmt.Printf("# cmd=%s\n", c.Command)
	return nil
}
