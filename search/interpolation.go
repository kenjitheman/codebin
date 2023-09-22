package search

func InterpolationSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right && target >= arr[left] && target <= arr[right] {
		if left == right {
			if arr[left] == target {
				return left
			}
			return -1
		}

		pos := left + ((right-left)/(arr[right]-arr[left]))*(target-arr[left])

		if arr[pos] == target {
			return pos
		}

		if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return -1
}
