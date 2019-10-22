package piscine

func CheckL(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsLower(str string) bool {
	runes := []rune(str)
	counter := 0
	for _, letter := range runes {
		if CheckL(letter) {
			counter++
		}
	}
	if counter == CheckStr(str) {
		return true
	}
	return false
}
