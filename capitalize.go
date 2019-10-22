package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - 32
	}
	for i := 1; i < Length(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' && ((runes[i-1] > '9' || runes[i-1] < '0') && runes[i] > 'z' || runes[i] < 'a' || runes[i] > 'Z' || runes[i] < 'A') {
			runes[i] -= 32
		} else if runes[i] >= 'A' && runes[i] <= 'Z' && ((runes[i-1] > '9' || runes[i-1] < '0') && runes[i] > 'Z' || runes[i] < 'A' || runes[i] > 'z' || runes[i] < 'a') {
			runes[i] += 32
		}

	}
	return string(runes)
}
