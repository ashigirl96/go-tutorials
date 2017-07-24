package moretypes

import (
	"fmt"
	"strings"
)

func SliceArray() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes[1:4])
}


func LiteralSliceArray() {
	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{1, false},
		{10, true},
	}
	fmt.Println(s[1])
}


func MakeSlice() {
	s := make([]int, 6, 10)  // len=6, cap=10
	fmt.Println(s)
}


func MakeBoard() {
	board := [][]string {
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}


func AppendArray() {
	var a []int
	b := []int{10, 20, 30}

	a = append(a, b)

	fmt.Println(a)
}

