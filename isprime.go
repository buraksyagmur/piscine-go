package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	} else {
		a := nb
		for i := 2; i < a; i++ {
			if a%i == 0 {
				return false
			}
		}
		return true
	}
}
