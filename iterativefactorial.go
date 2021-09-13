package piscine

func IterativeFactorial(nb int) int {
	a := nb - 1
	if nb < 1 {
		return 0
	} else if nb > 1 {
		for i := 1; i <= a; i++ {
			nb *= i
		}
	} else if nb == 1 {
		return 1
	}
	return nb
}
