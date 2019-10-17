package piscine

func UltimateDivMod(a *int, b *int) {

	division := *a / *b
	modulo := *a % *b
	*a = division
	*b = modulo
}
