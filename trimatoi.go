package piscine

func TrimAtoi(s string) int {
	a := []rune(s)
	b := 1
	c := 0
	neg := false
	for _, r := range string(a) {
		if !neg {
			if r == '-' {
				b = -1
			} else if r == '+' {
				b = 1
			}
		}
		if r >= '0' && r <= '9' {
			neg = true
			c = 10*c + int(r-'0')
		}
	}
	return b * c
}
