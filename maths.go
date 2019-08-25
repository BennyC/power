package power

// Min returns the minimum integer out of a and b
func Min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

// Avg returns the average value out of a given list of Points within a GPX file
func Avg(args ...Point) float64 {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i].Power
	}

	return float64(sum) / float64(len(args))
}
