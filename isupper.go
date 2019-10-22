package piscine

func CheckU(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func IsUpper(str string) bool {
	runes := []rune(str)
	counter := 0
	for _, letter := range runes {
		if CheckU(letter) {
			counter++
		}
	}
	if counter == CheckStr(str) {
		return true
	}
	return false
}
