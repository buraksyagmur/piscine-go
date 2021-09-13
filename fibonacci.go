package piscine

func Fibonacci(index int) int {
	if index > 1 {
		a := index - 1
		b := index - 2
		return Fibonacci(a) + Fibonacci(b)

	} else if index == 0 {
		return 0
	} else if index < 0 {
		return -1
	}
	return index
}
