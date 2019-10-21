package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb == 0 || nb == 1 {
		result = 1
		return result
	} else if nb < 25 && nb > 0 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		return result
	} else {
		result = 0
		return result
	}
}
