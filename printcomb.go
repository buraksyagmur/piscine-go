package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				if a < b && b < c && a < c && a != b && b != c && a != '7' {

					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else if a == '7' && b == '8' && c == '9' {
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					z01.PrintRune('\n')
				}
			}
		}
	}
}
