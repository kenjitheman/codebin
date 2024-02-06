package problems

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, e := range nums {
		complement := target - e
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[e] = i
	}
	return nil
}

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var target = 7

func TwoSumSorted(nums []int, target int) []int { 
	s := 0 
	e := len(nums) - 1
	for s < e {
		sum := nums[s] + nums[e]
		if sum == target {
			return []int{s, e}
		} else if sum < target {
			s++
		} else {
			e--
		}
	}
	return nil
}
