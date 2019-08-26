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
	Duration int   `json:"duration"`
	Power    Power `json:"power"`
}

// Power is a special type for the outputting of Power within a JSON marshalling
type Power float64

// MarshalJSON is custom marshalling for the Power output
// Changes the value to 2 dp
func (p Power) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", p)), nil
}

// Output the GPX power as JSON
func (j JSONOutputter) Output(g *power.Gpx, dur []int) {
	// build a slice of entries which will be provided to main output
	var ee []*Entry
	for _, i := range dur {
		ee = append(ee, &Entry{
			Duration: i * 60,
			Power:    Power(g.CalculateBestPower(i * 60)),
		})
	}

	// marshal the main wrapper which contains the gpx file name with all entries
	o := &Out{Name: g.Name, Entries: ee}
	bb, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		fmt.Println("something went wrong marshalling the results")
		return
	}

	fmt.Println(string(bb))
}
