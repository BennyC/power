package power

// Cli struct is the CLI application struct
type Cli struct {
}

// NewCli returns a new version of the CLI application
func NewCli() *Cli {
	return &Cli{}
}

// Run the Cli Application and return an exit code for the process to output
// to the terminal
func (c *Cli) Run() int {
	return 0
}
