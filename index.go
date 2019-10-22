package piscine

func Index(s string, toFind string) int {
	len := 0
	c := 0
	for i := range toFind {
		if i == i {
			len++
		}
	}
	for _, j := range toFind {
		for k, j1 := range s {
			if j == j1 {
				if len > 1 {
					for m := 0; m < len; m++ {
						if s[k+m] == toFind[m] {
							c++
						} else {
							return -1
						}
					}
					if c == len {
						return k
					}
				} else if len == 1 {
					return k
				} else {
					return -1
				}
			}
		}
		if c <= 0 {
			return -1
		}

	}
	return len
}
