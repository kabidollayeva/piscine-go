package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var i rune = '0'
	var j rune = '1'
	var k rune = '2'
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for kRune := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i < 7 {
					z01.PrintRune(44)
				} else {
					z01.PrintRune(' ')

				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	piscine.PrintComb()
}
