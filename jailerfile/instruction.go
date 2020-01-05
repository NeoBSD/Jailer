package jailerfile

// Instruction represents any Jailerfile instruction.
type Instruction interface {
	Execute() error
	Name() string
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

// CopyInstruction sets the working directory at build time inside the jail
type CopyInstruction struct {
	Command string `json:"copy"`
}

// Execute is run inside the jail
func (w CopyInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (w CopyInstruction) Name() string {
	return "COPY"
}

// CmdInstruction sets the working directory at build time inside the jail
type CmdInstruction struct {
	Command string `json:"cmd"`
}

// Execute is run inside the jail
func (w CmdInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (w CmdInstruction) Name() string {
	return "CMD"
}
