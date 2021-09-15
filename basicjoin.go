package piscine

func BasicJoin(elems []string) string {
	a := []string(elems)
	b := len(a)
	c := ""

	for i := 0; i < b; i++ {
		c += a[i]
	}
	return c
}
