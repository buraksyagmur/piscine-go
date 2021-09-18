package piscine

func SplitWhiteSpaces(s string) []string {
	result := ""
	var result2 []string
	a := len(s)
	srune := []rune(s)
	for i := 0; i < a; i++ {
		if srune[i] == 32 || srune[i] == '\t' || srune[i] == '\n' {
			if result != "" {
				result2 = append(result2, result)
				result = ""
			}
		} else {
			result += string(srune[i])
		}
		if i == a-1 {
			result2 = append(result2, result)
		}
	}
	return result2
}
