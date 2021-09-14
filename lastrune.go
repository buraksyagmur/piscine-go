package piscine

func LastRune(s string) rune {
	a := len(s)
	return []rune(s)[a-1]
}
