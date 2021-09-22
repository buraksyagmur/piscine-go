package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for k := i + 1; k < len(a); k++ {
			if a[i] == a[k] {
				a[i] = -1
				a[k] = -1
				break
			}
		}
	}
	for m := 0; m < len(a); m++ {
		if a[m] != -1 {
			return a[m]
		}
	}
	return -1
}
