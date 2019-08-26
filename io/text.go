package io

import (
	"fmt"

	"github.com/fudge/power"
)

// TextOutputter is for CLI usage
type TextOutputter struct {
}

// Output visualises the data for a Gpx struct in text format, used for
// CLI output
func (t TextOutputter) Output(g *power.Gpx, dur []int) {
	fmt.Println(g.Name)
	for _, i := range dur {
		p := g.CalculateBestPower(i * 60)
		fmt.Printf("-> best %v min power: %.2f\n", i, p)
	}
}
