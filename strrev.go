package piscine

func StrRev(s string) string {
	runes := []rune(s)
	for hello, world := 0, len(runes)-1; hello < world; hello, world = hello+1, world-1 {
		runes[hello], runes[world] = runes[world], runes[hello]
	}
	return string(runes)
}
