package jailerfile

// Instruction represents any Jailerfile instruction.
type Instruction interface {
	Execute() error
	Name() string
	Parse(source string) error
}

// FromInstruction specifies te base jail
type FromInstruction struct {
	From string `json:"from"`
}

// Execute is run inside the jail
func (f FromInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (f FromInstruction) Name() string {
	return "FROM"
}

// Parse parses the source string
func (f FromInstruction) Parse(input string) error {
	return nil
}

// RunInstruction executes a command at build time inside the jail
type RunInstruction struct {
	Command string `json:"run"`
}

// Execute is run inside the jail
func (r RunInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (r RunInstruction) Name() string {
	return "RUN"
}

// Parse parses the source string
func (r *RunInstruction) Parse(input string) error {
	r.Command = cleanString(input)
	return nil
}

// WorkDirInstruction sets the working directory at build time inside the jail
type WorkDirInstruction struct {
	Command string `json:"work_dir"`
}

// Execute is run inside the jail
func (w WorkDirInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (w WorkDirInstruction) Name() string {
	return "WORKDIR"
}

// Parse parses the source string
func (w *WorkDirInstruction) Parse(input string) error {
	w.Command = cleanString(input)
	return nil
}

// CopyInstruction sets the working directory at build time inside the jail
type CopyInstruction struct {
	Command string `json:"copy"`
}

// Execute is run inside the jail
func (c CopyInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (c CopyInstruction) Name() string {
	return "COPY"
}

// Parse parses the source string
func (c *CopyInstruction) Parse(input string) error {
	c.Command = cleanString(input)
	return nil
}

// CmdInstruction sets the working directory at build time inside the jail
type CmdInstruction struct {
	Command string `json:"cmd"`
}

// Execute is run inside the jail
func (c CmdInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (c CmdInstruction) Name() string {
	return "CMD"
}

// Parse parses the source string
func (c *CmdInstruction) Parse(input string) error {
	c.Command = cleanString(input)
	return nil
}
