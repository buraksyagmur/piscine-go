package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	name2 := []rune(name)
	a := len(name2)
	for i := 0; i < a; i++ {
		z01.PrintRune(name2[i])
	}
	z01.PrintRune('\n')
}
