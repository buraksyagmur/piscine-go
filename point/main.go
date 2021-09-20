package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{}

	setPoint(points)
	var aRune string = "x= "
	var bRune string = ", y= "
	z01.PrintRune(rune(aRune[0]))
	z01.PrintRune(rune(aRune[1]))
	z01.PrintRune(rune(aRune[2]))
	z01.PrintRune(rune(points.x[0]))
	z01.PrintRune(rune(points.x[1]))
	z01.PrintRune(rune(bRune[0]))
	z01.PrintRune(rune(bRune[1]))
	z01.PrintRune(rune(bRune[2]))
	z01.PrintRune(rune(bRune[3]))
	z01.PrintRune(rune(bRune[4]))
	z01.PrintRune(rune(points.y[0]))
	z01.PrintRune(rune(points.y[1]))
	z01.PrintRune('\n')
}
