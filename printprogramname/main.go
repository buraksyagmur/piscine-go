package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := []rune(os.Args[0])
	a := len(name)
	for i := 2; i < a; i++ {
		z01.PrintRune(name[i])
	}
	z01.PrintRune('\n')
}
