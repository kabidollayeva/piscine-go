package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	aString := []rune(str)
	for _, str := range aString {
		z01.PrintRune(str)
	}
	z01.PrintRune(10)
}
