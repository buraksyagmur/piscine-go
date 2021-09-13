package piscine

func IsPrime(nb int) bool {
	a := nb
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
