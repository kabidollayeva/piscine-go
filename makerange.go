package piscine

func MakeRange(min, max int) []int {
	size := max - min

	i := 0
	array := make([]int, i)
	if max < min || max == min {
		return array
	}
	answer := make([]int, size)
	for i := 0; i < size; i++ {
		answer[i] = i + min
	}
	return answer

}
