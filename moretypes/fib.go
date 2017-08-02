package moretypes

var table = make(map[int]int)

func Fib(n int) int {
	if n <
	if fn, ok := table[n]; ok {
		return fn
	}
	if n < 2 {
		return 1
	}
	fn := Fib(n-1) + Fib(n-2)
	table[n] = fn
	return fn
}