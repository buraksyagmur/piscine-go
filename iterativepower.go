package piscine

func IterativePower(nb int, power int) int {
	if power > 1 && power <= 30 {

		a := nb
		b := power

		for i := b; i >= 2; i-- {
			nb = a * nb
		}
	} else if power < 0 {
		return 0
	} else if power == 1 {
		return 1
	}
	return nb
}
