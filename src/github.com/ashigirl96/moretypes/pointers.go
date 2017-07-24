package moretypes

import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 10
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
