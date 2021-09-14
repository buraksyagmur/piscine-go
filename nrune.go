package piscine

func NRune(s string, n int) rune {
	if n > len(s) {
		return 0
	} else if n <= 0 {
		return 0
	} else if n >= 1 {
		return []rune(s)[n-1]
	}
	return 2
}
