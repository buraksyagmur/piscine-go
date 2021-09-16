package piscine

func MakeRange(min, max int) []int {
	if min > max || min == '0' && max == '0' {
		return nil
	}

	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}
