package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var a []int
	if n == 0 {
		a = append(a, 0)
	}
	for i := 0; n > 0; i++ {
		a = append(a, n%10)
		n /= 10
	}
	for index, num := range a {
		if num != 0 {
			for i := 0; i <= num; i++ {
				z01.PrintRune(rune(index + a[i]))
			}
		}
	}
}
