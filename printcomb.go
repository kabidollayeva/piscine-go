package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i <= 55; i++ {
		for j := i + 1; j <= 56; j++ {
			for k := j + 1; k <= 57; k++ {
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				if i < 55 {
					z01.PrintRune(',')
				} else {
					z01.PrintRune('\n')

				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	piscine.PrintComb()
}
