package sort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}

	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}
