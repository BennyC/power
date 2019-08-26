package cli

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fudge/power"
	"github.com/fudge/power/io"
)

// ErrNoFilePath when no file path has been provided via the args for Run
var ErrNoFilePath = errors.New("no file path provided")

// Cli struct is the CLI application struct
type Cli struct {
	Path        string
	Files       []*power.Gpx
	PowerOutput []int
	power.OutputType
}

// NewCli returns a new version of the CLI application
func NewCli() *Cli {
	// @todo need to flag out the output type
	return &Cli{
		OutputType:  power.OutputText,
		PowerOutput: []int{1, 5, 10, 20},
	}
}

// Run the Cli Application and return an exit code for the process to output
// to the terminal
func (c *Cli) Run(args ...string) (int, error) {
	if len(args) == 0 {
		return 2, ErrNoFilePath
	}

	if err := c.Load(args[0]); err != nil {
		fmt.Printf("err: %v", err)
		return 2, nil
	}

	for _, f := range c.Files {
		c.Output(f)
	}

	return 0, nil
}

// Load all GPX files at a given path, will also load a single GPX if the path
// is directly the GPX file
func (c *Cli) Load(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	switch m := f.Mode(); {
	case m.IsDir():
		pp, err := filepath.Glob(filepath.Clean(path) + "/*.gpx")
		if err != nil {
			return err
		}

		for _, p := range pp {
			c.File(p)
		}
	case m.IsRegular():
		c.File(path)
	}

	return nil
}

// File reads a file and load its contents in the Files slice within the CLI
func (c *Cli) File(p string) error {
	f, err := os.Open(p)
	if err != nil {
		return err
	}

	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	var gpx power.Gpx
	xml.Unmarshal(byteValue, &gpx)
	c.Files = append(c.Files, &gpx)
	return nil
}

// Output power readings for a GPX file
func (c *Cli) Output(g *power.Gpx) {
	var out power.Outputter
	switch c.OutputType {
	case power.OutputText:
		out = io.TextOutputter{}
	case power.OutputJSON:
		out = io.JSONOutputter{}
	}

	out.Output(g, c.PowerOutput)
}
