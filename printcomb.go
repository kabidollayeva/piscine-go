package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var i rune
	var j rune
	var k rune
	for i := 48; i <= 55; i++ {
		for j := i + 1; j <= 56; j++ {
			for k := j + 1; k <= 57; k++ {
				z01.PrintRune(i + 48)
				z01.PrintRune(j + 48)
				z01.PrintRune(k + 48)
				if i < 55 {
					z01.PrintRune(44)
				} else {
					z01.PrintRune(32)

				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	piscine.PrintComb()
}
