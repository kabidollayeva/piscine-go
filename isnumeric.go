package piscine

func CheckN(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func IsNumeric(str string) bool {
	runes := []rune(str)
	counter := 0
	for _, letter := range runes {
		if CheckN(letter) {
			counter++
		}
	}
	if counter == CheckStr(str) {
		return true
	}
	return false
}
