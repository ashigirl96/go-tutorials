package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "q" {
			break
		} else {
			counts[input.Text()]++
		}
	}
	// NOTE: ignoring potential erros from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%v: %s\n", n, line)
		}
	}

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
