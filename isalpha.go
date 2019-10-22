package piscine

func IsAlpha(str string) bool {
	runes := []rune(str)
	counter := 0
	for _, letter := range runes {
		if Check(letter) {
			counter++
		}
	}
	if counter == CheckStr(str) {
		return true
	}
	return false
}

func Check(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}

func CheckStr(str string) int {
	var count int
	aRune := []rune(str)
	for index := range aRune {
		count = index + 1
	}
	return count
}
