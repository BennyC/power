package power

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
