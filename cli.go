package power

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ErrNoFilePath when no file path has been provided via the args for Run
var ErrNoFilePath = errors.New("no file path provided")

// ErrFileDoesNotExist when no file exists at a given path
var ErrFileDoesNotExist = errors.New("file or directory does not exist")

// Cli struct is the CLI application struct
type Cli struct {
	Path  string
	Files []*Gpx
}

// NewCli returns a new version of the CLI application
func NewCli() *Cli {
	return &Cli{}
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

	var gpx Gpx
	xml.Unmarshal(byteValue, &gpx)
	c.Files = append(c.Files, &gpx)
	return nil
}

// Output power readings for a GPX file
func (c *Cli) Output(g *Gpx) {
	fmt.Println(g.Name)
	for _, i := range [4]int{1, 5, 10, 20} {
		p := g.CalculateBestPower(i * 60)
		fmt.Printf("-> best %v min power: %.2f\n", i, p)
	}
}
