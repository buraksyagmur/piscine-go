package piscine

func ActiveBits(n int) uint {
	d := 2
	if n%2 == 0 {
		d = 1
	}
	for n != 1 {
		if n == 1 {
			break
		} else {
			n = n / 2
			d += n % 2
		}
	}
	return uint(d-1)
}
