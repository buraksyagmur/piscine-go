package piscine

func Map(f func(int) bool, a []int) []bool {
	var c []bool
	for i := 0; i < len(a); i++ {
		c = append(c, f(a[i]))
	}
	return c
}
