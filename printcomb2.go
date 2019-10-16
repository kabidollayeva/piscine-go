package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= 9; j++ {
			k := j + 1
			for n := i; n <= '9'; n = n + 1 {
				for ; k <= '9'; k = k + 1 {
					z01.PrintRune((i))
					z01.PrintRune((j))
					z01.PrintRune((k))
					z01.PrintRune(' ')
					z01.PrintRune((n))
					if i < '9' || j < '8' || n < '9' || k < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				k = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
