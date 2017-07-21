package conditions

import (
	"math"
)

func Pow(x, n, lim float64) float64 {
	if z := math.Pow(x, n); z < lim {
		return z
	}
	return lim
}
