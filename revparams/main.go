package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(arg) - 1; i >= 1; i-- {
		str := []string(arg)
		runes := []rune(str[i])
		for _, word := range runes {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
