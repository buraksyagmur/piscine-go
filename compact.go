package piscine

func Compact(ptr *[]string) int {
	result := 0
	var result1 []string
	for _, r := range *ptr {
		if r != "" {
			result1 = append(result1, r)
			result++
		}
	}
	*ptr = result1
	return result
}
