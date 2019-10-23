package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	str := []string(arg)
	for index, v := range str {
		if index != 0 {
			runes := []rune(v)
			for _, runes := range runes {
				z01.PrintRune(runes)
			}
			z01.PrintRune('\n')
		}
	}
}
