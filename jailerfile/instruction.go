package jailerfile

// Instruction represents any Jailerfile instruction.
type Instruction interface {
	Execute() error
	Name() string
}

// FromInstruction specifies te base jail
type FromInstruction struct {
	From string
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
	Command string
}

// Execute is run inside the jail
func (r RunInstruction) Execute() error {
	return nil
}

// Name returns the instruction identifier
func (r RunInstruction) Name() string {
	return "RUN"
}
