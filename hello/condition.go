package main

import (
	"fmt"
	"github.com/ashigirl96/conditions"
)

func main() {
	fmt.Println(
		conditions.Pow(3, 2, 10),
		conditions.Pow(3, 3, 20),
	)

	fmt.Println()

	fmt.Printf(
		"Newton-Sqrt %g,\nCommon-Sqrt %g\n",
		conditions.NewTonSqrt(2), conditions.Sqrt2(2),
	)

	fmt.Println()

	conditions.SwitchMain()

	fmt.Println()

	conditions.SwitchEvaluationOrder()

	fmt.Println()

	conditions.DeferMain()
}
