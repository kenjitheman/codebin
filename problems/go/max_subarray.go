package problems

// FindMaxSubarray returns the maximum subarray of a given array

var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func FindMaxSubarray(arr []int) []int {
    var maxSum int
    var currentSum int
    var maxStartIndex int
    var maxEndIndex int
    var currentStartIndex int
    var currentEndIndex int

    for i := 0; i < len(arr); i++ {
        if currentSum <= 0 {
            currentStartIndex = i
            currentSum = arr[i]
        } else {
            currentSum += arr[i]
        }

        if currentSum > maxSum {
            maxSum = currentSum
            maxStartIndex = currentStartIndex
            maxEndIndex = i
        }
    }

    return arr[maxStartIndex : maxEndIndex+1]
}
