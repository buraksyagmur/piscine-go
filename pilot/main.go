package main

import (
	"fmt"
)

type Pilot struct {
	Name     string
	Life     int
	Age      int
	Aircraft string
}

func main() {
	AIRCRAFT1 := "1"
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}
