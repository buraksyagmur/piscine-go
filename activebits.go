package piscine

func ActiveBits(n int) int {
	d := 1
	if n%2 == 0 {
		d = 0
	}
	for n != 1 {
		if n == 1 {
			break
		} else {
			n = n / 2
			d += n % 2
		}
	}
	return int(d)
}
