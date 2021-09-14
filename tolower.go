package piscine

func ToLower(s string) string {
	a := []rune(s)
	b := len(a) - 1
	c := ""
	for i := 0; i <= b; i++ {
		if a[i] >= 65 && a[i] <= 90 {
			a[i] = a[i] + 32
		}
		c += string(a[i])

	}
	return c
}
