package conditions

import (
	"math"
	"fmt"
)

func Pow(x, n, lim float64) float64 {
	if z := math.Pow(x, n); z < lim {
		return z
	} else {
		fmt.Printf("%g >= %g", z, lim)
	}
	return lim
}
