package power

import (
	"errors"
	"fmt"
	"strconv"
)

// Output types which are configured via flags
const (
	OutputText = iota
	OutputJSON
	OutputSvg
)

// Outputter is an interface for outputting an array of GPX structs
// into the power formats required
type Outputter interface {
	Output(g *Gpx, dur []int)
}

// OutputType which will hold flag info
type OutputType int

// Set provides the flag setting information needed for the CLI application
func (d *OutputType) Set(value string) error {
	v, err := strconv.Atoi(value)

	if v > OutputSvg || v < OutputText || err != nil {
		return errors.New("not a valid output type")
	}

	*d = OutputType(v)
	return nil
}

// String provides the string representation of the value for the CLI application
func (d *OutputType) String() string {
	return fmt.Sprint(*d)
}
