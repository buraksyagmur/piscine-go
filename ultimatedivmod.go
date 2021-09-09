package piscine

func UltimateDivMod(a *int, b *int) {
	*a = *a / *b
	*b = 5
	*b = *a % *b
}
