package piscine

func Index(s string, toFind string) int {
	c := len(s)
	d := len(toFind)

	for i := 0; i < c-d; i++ {
		if toFind == s[i:i+d] {
			return i
		}
	}
	return -1
}
