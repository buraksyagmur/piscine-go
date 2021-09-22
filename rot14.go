package piscine

func Rot14(s string) string {
	c := []rune(s)
	result := ""
	for i := 0; i < len(s); i++ {
		if c[i] >= 'A' && c[i] < 'M' || c[i] >= 'a' && c[i] < 'm' {
			c[i] += 14
		} else if c[i] >= 'm' && c[i] <= 'z' || c[i] >= 'M' && c[i] <= 'Z' {
			c[i] -= 12
		}
		result += string(c[i])
	}
	return result
}
