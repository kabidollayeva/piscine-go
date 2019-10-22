package piscine

func Length(s string) int {
	counter := 0
	for i := range s {
		if i == i {
			counter++
		}
	}
	return counter
}
func NRune(s string, n int) rune {
	sentence := []rune(s)
	len := Length(s)
	if n <= len && n-1 >= 0 {
		return sentence[n-1]
	} else {
		return 0
	}
}
