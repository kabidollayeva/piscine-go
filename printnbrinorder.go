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
	for i := range a {
		z01.PrintRune(rune('0' + a[i]))
	}
}
