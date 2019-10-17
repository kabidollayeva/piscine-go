package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string{
	aString := []rune(s)
	for i, j := 0; len(aString)-1; i < j; i, j = i + 1, j - 1 {
		aString[i], aString[j] = aString[j], aString[i]
	}
	return string(aString)
}
