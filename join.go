package piscine

func Join(strs []string, sep string) string {
	a := []string(strs)
	b := len(a)
	c := ""

	for i := 0; i < b; i++ {
		if a[i] == "?" {
			c += a[i]
		}
		if i == (b - 1) {
			c += a[i]
		} else {
			c += a[i] + ":"
		}
	}
	return c
}
