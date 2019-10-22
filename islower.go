package piscine

func IsLower(str string) bool {
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
