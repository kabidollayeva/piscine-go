package piscine

func StrRev(s string) string {
	aString := []rune(s)
	var len int = 0
	var str rune
	for v := range aString {
		len++
		v = v
	}
	for i := 0; i < len/2; i++ {
		str = aString[i]
		aString[i] = aString[len-i-1]
		aString[len-i-1] = str
	}
	return string(aString)
}
