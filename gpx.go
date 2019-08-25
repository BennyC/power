package power

import (
	"encoding/xml"
	"math"
	"time"
)

// Gpx represents a gpx file
type Gpx struct {
	XMLName  xml.Name  `xml:"gpx"`
	Name     string    `xml:"trk>name"`
	Segments []Segment `xml:"trk>trkseg"`
}

// CalculateBestPower will take a duration in seconds as a parameter
// and provide the best average power for this duration within the GPX file
func (g Gpx) CalculateBestPower(secs int) float64 {
	var max float64
	var pp []Point
	for _, s := range g.Segments {
		pp = append(pp, s.Points...)
	}

	for i := 0; i < len(pp)-secs; i++ {
		j := Min(i+secs, len(pp))
		avg := Avg(pp[i:j]...)
		max = math.Max(max, avg)
	}

	return max
}

// Segment represents a "trkseg" within a Gpx file, commonly refered to as a "lap"
type Segment struct {
	XMLName xml.Name `xml:"trkseg"`
	Points  []Point  `xml:"trkpt"`
}

// Point represents a portion of time within a Segment, generally contains power
// and lat long information
type Point struct {
	XMLName xml.Name  `xml:"trkpt"`
	Time    time.Time `xml:"time"`
	Power   int       `xml:"extensions>power"`
}
