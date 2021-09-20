package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if (os.Args[1]) == "quest8.txt" {
		file, err := os.Open("quest8.txt")
		if err != nil {
			fmt.Printf("the mistake is: %v\n", err.Error())
		}
		a := make([]byte, 26)
		file.Read(a)
		fmt.Print(string(a))
	}
}
