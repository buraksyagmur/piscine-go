package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	} else {
		for i := 1; i <= nb; i++ {
			if !IsPrime(nb) {
				nb = nb + i
			} else {
				return nb
			}
		}
	}
	return nb
}
