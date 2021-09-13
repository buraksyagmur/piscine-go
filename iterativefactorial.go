package piscine

func IterativeFactorial(nb int) int {
	if nb <= 1 {
		return 0
	} else if nb > 1 {
		a := nb - 1
		for i := 1; i <= a; i++ {
			nb *= i
		}
	}
	return nb
}
