package moretypes

func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


//func Fib() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a + b
//		return a
//	}
//}
