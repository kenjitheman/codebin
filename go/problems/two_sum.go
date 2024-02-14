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

func TwoSumSorted(nums []int, target int) []int {
    s, e := 0, len(nums)-1
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
