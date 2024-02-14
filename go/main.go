package main

import "github.com/plastiey/codegen/problems"

func main() {
	var nums = []int{2, 7, 7, 1, 3, 3, 2, 11, 15, 8, 4, 9, 6, 5, 10, 12, 13, 14}
	var target = 12
	problems.TwoSum(nums, target)
}
