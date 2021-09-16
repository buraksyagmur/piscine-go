package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	}

	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}
