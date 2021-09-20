package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	result := ""
	b := len(a)
	for i := 0; i < b; i++ {

		if i != 0 && string(a[i]) != " " {
			result += "\n"
		}
		{
			result += string(a[i])
		}
	}
	z01.PrintRune(rune(result[0]))
	z01.PrintRune(rune(result[1]))
	z01.PrintRune(rune(result[2]))
	z01.PrintRune(rune(result[3]))
	z01.PrintRune(rune(result[4]))
	z01.PrintRune(rune(result[5]))
	z01.PrintRune(rune(result[6]))
	z01.PrintRune(rune(result[7]))
	z01.PrintRune(rune(result[8]))
	z01.PrintRune(rune(result[9]))
	z01.PrintRune(rune(result[10]))
	z01.PrintRune(rune(result[11]))
	z01.PrintRune(rune(result[12]))
	z01.PrintRune(rune(result[13]))
	z01.PrintRune(rune(result[14]))
	z01.PrintRune(rune(result[15]))
	z01.PrintRune(rune(result[16]))
	z01.PrintRune(rune(result[17]))
}
