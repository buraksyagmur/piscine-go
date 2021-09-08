package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var aRune string = "789"
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				for d := ','; d <= ','; d++ {
					for e := ' '; e <= ' '; e++ {

						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune(d)
						z01.PrintRune(e)
					}
				}
			}
		}
	}
	z01.PrintRune(rune(aRune[0]))
	z01.PrintRune(rune(aRune[1]))
	z01.PrintRune(rune(aRune[2]))
	z01.PrintRune('\n')
}
