package piscine

func IsNumeric(s string) bool {
	a := []rune(s)
	b := len(a) - 1
	for i := 0; i <= b; i++ {
		if (a[i] >= 0) && (a[i] <= 47) || a[i] >= 58 && a[i] <= 127 {
			return false
		}
	}
	return true
}
