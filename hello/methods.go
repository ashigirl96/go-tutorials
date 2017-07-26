package main

import (
	"fmt"
	"github.com/ashigirl96/methods"
	"math"
	"golang.org/x/tour/pic"
)

func describe(i methods.I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


func main() {
	var f methods.MyFloat64
	sqrt2 := methods.MyFloat64(math.Sqrt2)
	f = -20
	v := methods.Vertex{3, 4}
	v2 := &methods.Vertex{3, 4}

	var i methods.I


	fmt.Println(v.Abs2())
	fmt.Println(methods.Abs2(v))
	fmt.Println("Before Scale:", v)
	v.Scale(100)
	methods.ScaleFunc(v2, 100)
	fmt.Println("After Scale(100): ", v)
	fmt.Println("After ScaleFunc: ", v2)
	v2.Scale(3)
	fmt.Println("After ScaleFunc: ", v2)

	fmt.Println(f.Abs())
	fmt.Println(sqrt2.Abs())
	var a methods.Abser

	fmt.Println()

	a = f
	fmt.Println("if a == Float64, a.Abs() ==> ", a.Abs())

	a = &v
	fmt.Println("if a == Vertex, a.Abs() ==> ", a.Abs())

	fmt.Println()

	i = f
	describe(i)
	f.M() // -20

	i = &methods.T{"Hello"}
	describe(i)
	i.M() // Hello

	fmt.Println()

	var ii interface{}

	methods.Describe(ii)

	ii = 42

	methods.Describe(ii)

	ii = "Hello, 世界"

	methods.Describe(ii)

	fmt.Println()
	fmt.Println()

	s, ok := ii.(string)
	fmt.Println(s, ok)

	ff, ok := ii.(float64)
	fmt.Println(ff, ok)

	fmt.Println()

	methods.TypeSwitchesDo(21)
	methods.TypeSwitchesDo("Hello, World")
	methods.TypeSwitchesDo(true)

	fmt.Println()
	fmt.Println()

	arthur := methods.Person{"Aruthur Dent", 42}
	zaphod := methods.Person{"Zaphod Beeblebrox", 90001}
	fmt.Println(arthur, "\n", zaphod)


	fmt.Println()
	fmt.Println()

	hosts := map[string]methods.IPAddr {
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v, %v\n", name, ip)
	}

	if err := methods.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println()

	fmt.Println(methods.Sqrt(2))
	fmt.Println(methods.Sqrt(-2))

	fmt.Println()
	fmt.Println()

	// strings.Reader を作成し、 8byte毎に読み出し
	methods.ReaderErrorTest()

	fmt.Println()
	fmt.Println()

	img := methods.MyImage{32, 32}
	pic.ShowImage(img)
}

