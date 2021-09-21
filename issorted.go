package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	y := true
	w := true

	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			y = false
		}
	}

	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			w = false
		}
	}
	return w || y
}

func isSuperior(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func lenghtof(d []int) int {
	i := 0
	for range d {
		i++
	}
	return i
}
