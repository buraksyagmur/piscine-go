package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1:]
	lenA := len(os.Args[1:])

	for i := 0; i < lenA; i++ {
		if a[i] == "01" || a[i] == "galaxy" || a[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		} else {
			// fmt.Println(nil)
		}
	}
}
