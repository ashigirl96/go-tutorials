package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}


func times(x, y int) int {
	return x * y
}


func swap(x, y string) (string, string) {
	return y, x
}


func split(sum int) (x, y int) {
	x = sum * 4 / 7
	y = sum - x
	return
}


func main() {
	a, b := swap(", World", "Hello")
	fmt.Println("10 + 2 = ", add(10, 2))
	fmt.Println("5 * 4 = ", times(5, 4))
	fmt.Println(a, b)
	fmt.Println(split(17))
}
