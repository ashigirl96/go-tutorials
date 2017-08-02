package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	CAD
	JPY
	EUR
)

// like python's enum.


func main() {
	symbol := [...]string{USD: "$", JPY: "yen", CAD: "£", EUR: "¡"}  // TEKITOU

	fmt.Println(EUR, symbol[EUR])

	r := [...]int{10: -1}
	fmt.Println(r)
	a := [...]int{1, 2}
	fmt.Println(a)
}
