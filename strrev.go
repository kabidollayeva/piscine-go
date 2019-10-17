package piscine

func StrRev(s string) string {
	var reverse string
	for i := length(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}