package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var b [10]int
	for n != 0 {

		b[n%10]++
		n /= 10
	}
	for i := 0; i <= 9; i++ {
		for b[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			b[i]--
		}
	}
}
