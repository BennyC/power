package io

import (
	"github.com/fudge/power"
)

// SVGOutputter to output SVG to the CLI, which can then be piped by the User
// to a desired file
type SVGOutputter struct {
}

// Output the GPX file in an SVG format, which will be a chart
func (s SVGOutputter) Output(g *power.Gpx, dur []int) {
}
