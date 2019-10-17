package piscine

func StrLen(str string) int {
	count := 0
	for _, aString := range str {
		count++
		_ = aString
	}
	return count
}
