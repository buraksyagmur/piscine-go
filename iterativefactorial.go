package piscine

func IterativeFactorial(nb int) int {
	a := nb - 1
	if nb < 0 {
		return 0
	} else if nb > 1 && nb <= 40 {
		for i := 1; i <= a; i++ {
			nb *= i
		}
	} else if nb == 1 {
		return 1
	} else if nb >= 41 {
		return 0
	} else if nb == 0 {
		return 1
	}
	return nb
}
