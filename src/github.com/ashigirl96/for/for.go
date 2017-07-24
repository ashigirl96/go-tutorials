package main

import (
	"fmt"
)

func Pow() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2 ** %d == %d\n", i, v)
	}
}

func Pow2() {
	pow := make([]int, 8)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Println(value)
	}
}




func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	Pow()

	Pow2()

	var pic [8][3]uint8


	fmt.Println(pic)
	fmt.Println(len(pic[0]))
	fmt.Println(pic2)
}
