package main

import (
	//	"io"
	"os"
	//	"strings"
	"io/ioutil"

	"github.com/01-edu/z01"
)

func main() {
	var aRune string = "Hello"
	var bRune string = "^C"

	if len(os.Args) > 2 && (os.Args[1]) == "quest8.txt" && (os.Args[2]) == "quest8T.txt" {
		content, err := ioutil.ReadFile("quest8.txt")
		content2, err := ioutil.ReadFile("quest8T.txt")
		if err != nil {
			z01.PrintRune(rune(err.Error()[0]))
		}
		for i := 0; i < len(string(content)); i++ {
			z01.PrintRune(rune(content[i]))
		}
		for k := 0; k < len(string(content2)); k++ {
			z01.PrintRune(rune(content2[k]))
		}

	} else if len(os.Args) == 1 && (os.Args[1]) == "quest8.txt" {
		content, err := ioutil.ReadFile("quest8.txt")
		if err != nil {
			z01.PrintRune(rune(err.Error()[0]))
		}
		for i := 0; i < len(string(content)); i++ {
			z01.PrintRune(rune(content[i]))
		}

	} else if len(os.Args) == 1 && (os.Args[1]) == "quest8T.txt" {
		content, err := ioutil.ReadFile("quest8T.txt")
		if err != nil {
			z01.PrintRune(rune(err.Error()[0]))
		}
		for i := 0; i < len(string(content)); i++ {
			z01.PrintRune(rune(content[i]))
		}

	} else if len(os.Args) > 1 && (os.Args[1]) == "quest8.txt" || (os.Args[1]) != "quest8T.txt" {
		content, err := ioutil.ReadFile("quest8.txt")
		if err != nil {
			z01.PrintRune(rune(err.Error()[0]))
		}
		for i := 0; i < len(string(content)); i++ {
			z01.PrintRune(rune(content[i]))
		}
		// z01.PrintRune('\n')
		var cRune string = "ERROR: "
		var dRune string = os.Args[1]
		var eRune string = ": No such file or directory"

		z01.PrintRune(rune(cRune[0]))
		z01.PrintRune(rune(cRune[1]))
		z01.PrintRune(rune(cRune[2]))
		z01.PrintRune(rune(cRune[3]))
		z01.PrintRune(rune(cRune[4]))
		z01.PrintRune(rune(cRune[5]))
		z01.PrintRune(rune(cRune[6]))
		for i := 0; i < len(os.Args[1]); i++ {
			z01.PrintRune(rune(dRune[i]))
		}
		for k := 0; k < 27; k++ {
			z01.PrintRune(rune(eRune[k]))
		}
		z01.PrintRune('\n')

	} else if len(os.Args) > 1 && (os.Args[1]) != "quest8.txt" || (os.Args[1]) != "quest8T.txt" {
		var cRune string = "ERROR: "
		var dRune string = os.Args[1]
		var eRune string = ": No such file or directory"

		z01.PrintRune(rune(cRune[0]))
		z01.PrintRune(rune(cRune[1]))
		z01.PrintRune(rune(cRune[2]))
		z01.PrintRune(rune(cRune[3]))
		z01.PrintRune(rune(cRune[4]))
		z01.PrintRune(rune(cRune[5]))
		z01.PrintRune(rune(cRune[6]))
		for i := 0; i < len(os.Args[1]); i++ {
			z01.PrintRune(rune(dRune[i]))
		}
		for k := 0; k < 27; k++ {
			z01.PrintRune(rune(eRune[k]))
		}
		z01.PrintRune('\n')

	} else if len(os.Args) == 1 {
		z01.PrintRune(rune(aRune[0]))
		z01.PrintRune(rune(aRune[1]))
		z01.PrintRune(rune(aRune[2]))
		z01.PrintRune(rune(aRune[3]))
		z01.PrintRune(rune(aRune[4]))
		z01.PrintRune('\n')
		z01.PrintRune(rune(aRune[0]))
		z01.PrintRune(rune(aRune[1]))
		z01.PrintRune(rune(aRune[2]))
		z01.PrintRune(rune(aRune[3]))
		z01.PrintRune(rune(aRune[4]))
		z01.PrintRune('\n')
		z01.PrintRune(rune(bRune[0]))
		z01.PrintRune(rune(bRune[1]))
		z01.PrintRune('\n')
	}
}
