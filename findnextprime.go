package piscine

func FindNextPrime(nb int) int {
	a := nb
	if nb < 2 || nb > 1000000000000 {
		return 2
	} else {
		for i := 1; i < nb; i++ {
			if !IsPrime(nb) {
				nb = a + i
			} else {
				break
			}
		}
	}
	return nb
}
