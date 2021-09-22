package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	result := 0
	for start != 1 {

		if start%2 == 0 {
			start = start / 2
		} else if start%2 != 0 {
			start = (start * 3) + 1
		}
		result++
	}
	return result
}
