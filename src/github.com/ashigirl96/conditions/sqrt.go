package conditions

import (
	"fmt"
	"github.com/ashigirl96/newmath"
	"math"
)

func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt2(x float64) float64 {
	return newmath.Sqrt(x)
}
