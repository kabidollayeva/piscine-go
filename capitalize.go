package piscine

func Check(r rune) bool {
	if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
		return true

	}
	return false
}
func Capitalize(s string) string {
	runes := []rune(s)
	for index, letter := range runes {
		if Check(letter) {
			if index == 0 || Check(runes[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					runes[index] -= 32
				} else {
					if letter >= 'A' && letter <= 'Z' {
						runes[index] += 32
					}
				}
			}
		}

	}
	return string(runes)
}
