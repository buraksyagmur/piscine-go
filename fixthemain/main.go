package piscine

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	CLOSE = false
	OPEN  = true
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}
