package main

import (
	"fmt"
	"github.com/ashigirl96/moretypes"
	"golang.org/x/tour/wc"
)

var (
	v3 = moretypes.Vertex{1, 2}
	v4 = moretypes.Vertex{X: 1}
	p = &moretypes.Vertex{1, 2}
)

func main() {
	moretypes.Pointers()

	v1 := moretypes.Vertex{1, 2}
	v2 := moretypes.Vertex{3, 4}
	v2.Y = 1000
	fmt.Println(v1.X+v2.X, v1.Y+v2.Y)

	fmt.Println(v3, v4, p)

	moretypes.SliceArray()
	moretypes.LiteralSliceArray()

	s := [3]int{3, 2, 1}

	fmt.Println(len(s[0:1]))  // 1
	fmt.Println(cap(s[0:1]))  // 3

	var t []int

	if t == nil {
		fmt.Println("s is nil")
	}

	moretypes.MakeSlice()

	moretypes.MakeBoard()

	moretypes.Maps()

	wc.Test(moretypes.WordCount)

	su := moretypes.Adder()

	fmt.Println(su(10))
	fmt.Println(su(200))

	fib := moretypes.Fib()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
