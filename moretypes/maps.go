package moretypes

import (
	"fmt"
	"strings"
)

var (
	v int
	ok bool
)

func Maps() {
	m := make(map[string]Vertex)

	m["Bell"] = Vertex{10.3, 20.2}

	n := map[string]Vertex {
		"Bell": Vertex{1., 2.},
		"Call": Vertex{3., 4.},
	}

	b := map[string]Vertex{
		"A": {1., 2.},
		"B": {3., 4.},
	}

	c := make(map[string]int)

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(b)

	c["A"] = 10
	v, ok = c["A"]
	fmt.Println(v, ok)

	delete(c, "A")
	v, ok = c["A"]
	fmt.Println(v, ok)

	c["A"] += 1
	fmt.Println(c["A"])
}

func WordCount(s string) map[string]int {
	dic := make(map[string]int)

	for _, val := range strings.Fields(s) {
		dic[val] += 1
	}

	return dic
}
