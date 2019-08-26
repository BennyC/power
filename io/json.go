package io

import (
	"encoding/json"
	"fmt"

	"github.com/fudge/power"
)

// JSONOutputter outs JSON to the CLI instead of a basic human
// readable text output
type JSONOutputter struct {
}

// Out is the main output wrapper for the JSON output
type Out struct {
	Name    string   `json:"name"`
	Entries []*Entry `json:"entries"`
}

// Entry is an individual Power representation before going to CLI
type Entry struct {
	Duration int     `json:"duration"`
	Power    float64 `json:"power"`
}

// Output the GPX power as JSON
func (j JSONOutputter) Output(g *power.Gpx, dur []int) {
	var ee []*Entry
	for _, i := range dur {
		p := g.CalculateBestPower(i * 60)
		e := &Entry{
			Duration: i * 60,
			Power:    p,
		}
		ee = append(ee, e)
	}

	o := &Out{
		Name:    g.Name,
		Entries: ee,
	}

	bb, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		fmt.Println("something went wrong marshalling the results")
		return
	}

	fmt.Println(string(bb))
}
