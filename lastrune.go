package piscine

func LastRune(s string) rune {
	len := Length(s)
	return NRune(s, len)
}
