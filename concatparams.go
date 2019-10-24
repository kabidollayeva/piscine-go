package piscine

func ConcatParams(args []string) string {
	//	size := len(args)
	//	answer := make([]int, size)
	//count:=0
	//for range
	var ans string
	for index, word := range args {
		if index > 0 {
			ans = ans + "\n"
		}
		ans = ans + word

	}
	return ans
}
