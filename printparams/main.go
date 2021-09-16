package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args

	c := len(a)

	for i := 1; i < c; i++ {
		result := []rune(a[i])
		for k := 0; k < len(result); k++ {
			z01.PrintRune(result[k])
		}
		z01.PrintRune('\n')
	}
}
