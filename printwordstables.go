package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var result string
	b := len(a)
	for i := 0; i < b; i++ {
		result += a[i]
		if i != b {
			result += "\n"
		}
	}
	for _, str := range result {
		z01.PrintRune(str)
	}
}
