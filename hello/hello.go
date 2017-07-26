package main

import (
	"fmt"

	"github.com/ashigirl96/newmath"

	"github.com/ashigirl96/conditions"
)

func main() {
	fmt.Printf("Hello world, Sqrt(2) = %v\n", newmath.Sqrt(2))
	fmt.Printf("Hello world, Sqrt(2) = %v\n", conditions.Sqrt2(2))
}
