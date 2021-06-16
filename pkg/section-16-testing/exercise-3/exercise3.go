package myaverage

import (
	"sort"
)

// CenteredAvg computes the average of a list of numbers
// after removing the largest and smallest values.
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	x := xi[1:(len(xi) - 1)]
	s := 0
	for _, v := range x {
		s += v
	}
	return float64(s) / float64(len(x))
}
