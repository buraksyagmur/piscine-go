package piscine

func IsPrintable(s string) bool {
	a := []rune(s)
	b := len(a) - 1
	for i := 0; i <= b; i++ {
		if (a[i] >= 0) && (a[i] <= 31) {
			return false
		}
	}
	return true
}
