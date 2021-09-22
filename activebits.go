package piscine

func ActiveBits(n int) uint {
	counter := 0
	if n < 0 {
		n = -n
	}
	a := 0
	b := 0
	DivMod(n, 2, &b, &a)
	for b > 0 {
		DivMod(n, 2, &b, &a)
		n = b
		if a == 1 {
			counter++
		}
	}
	return uint(counter)
}
