package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a = a + 1 {
		for b := '0'; b <= '9'; b = b + 1 {
			for c := '0'; c <= '9'; c = c + 1 {
				for d := '0'; d <= '9'; d = d + 1 {

					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(' ')
					z01.PrintRune(rune(c))
					z01.PrintRune(rune(d))
					z01.PrintRune(',')
					z01.PrintRune(' ')

					if a == '9' && b == '9' && c == '9' && d == '9' {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(' ')
						z01.PrintRune(rune(c))
						z01.PrintRune(rune(d))

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
