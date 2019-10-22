package piscine

func Sing(s string) bool {
	S := false
	for _, v := range s {
		if v >= '0' && v <= '9' {
			break
		}
		if v == '-' {
			S = true
		}
	}
	return S
}

func TrimAtoi(s string) int {
	res := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			c := 0
			for i := '1'; i <= v; i++ {
				c++
			}
			res = res*10 + c
		}
	}
	if Sing(s) {
		res *= -1
	}
	return res
}
