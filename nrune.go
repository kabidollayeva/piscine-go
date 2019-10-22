package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)
	if n <= 9 {
		return sentence[n-1]
	} else {
		return 0
	}
}
