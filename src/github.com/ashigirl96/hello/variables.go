package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java, cpp bool
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "No!!"
	x, y := 3, 5
	const f = "%T(%v)\n"

	fmt.Println(i, j, c, python, java, cpp)
	fmt.Println(x, y)

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	t := math.Sqrt(float64(x*x + y*y))
	var z uint = uint(t)
	fmt.Println(x, y, t, z)

	const World = "5000兆円"

	fmt.Println(World, "ほしい")
}
