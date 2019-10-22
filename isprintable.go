package piscine

func CheckP(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}

func IsPrintable(str string) bool {
	runes := []rune(str)
	counter := 0
	for _, letter := range runes {
		if CheckP(letter) {
			return false
		}
	return true
}