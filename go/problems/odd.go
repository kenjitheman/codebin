package problems

func FindOdd(seq []int) int {
	res := 0
	for _, num := range seq {
		res ^= num
	}
	return res
}
