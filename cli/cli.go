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
	// TODO need to flag out the output type
	return &Cli{
		OutputType:  power.OutputText,
		PowerOutput: []int{1, 5, 10, 20},
	}
}

// Run the Cli Application and return an exit code for the process to output
// to the terminal
func (c *Cli) Run(args ...string) (int, error) {
	// TODO better error handling/reporting from calls
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

	var paths []string
	switch m := f.Mode(); {
	case m.IsDir():
		pp, _ := filepath.Glob(filepath.Clean(path) + "/*.gpx")
		paths = pp
	case m.IsRegular():
		paths = append(paths, path)
	}

	ch := make(chan power.Gpx, len(paths))
	loadFiles(paths, ch)

	for v := range ch {
		c.Files = append(c.Files, &v)
	}

	return nil
}

// loadFiles loads all paths given and passes the power.Gpx struct back across
// a channel, also handles the closing of the channel once done
func loadFiles(pp []string, ch chan power.Gpx) {
	for _, p := range pp {
		f, _ := os.Open(p)
		defer f.Close()

		byteValue, _ := ioutil.ReadAll(f)

		var gpx power.Gpx
		xml.Unmarshal(byteValue, &gpx)
		ch <- gpx
	}

	close(ch)
}

// Output power readings for a GPX file
func (c *Cli) Output(g *power.Gpx) {
	var out power.Outputter
	// TODO maybe put in slice & need to implement SVG
	switch c.OutputType {
	case power.OutputText:
		out = io.TextOutputter{}
	case power.OutputJSON:
		out = io.JSONOutputter{}
	}

	out.Output(g, c.PowerOutput)
}
