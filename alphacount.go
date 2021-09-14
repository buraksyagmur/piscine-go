package piscine

func AlphaCount(s string) int {
	n := 0
	for _, i := range s {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
			n++
		}
	}
	return n
}
