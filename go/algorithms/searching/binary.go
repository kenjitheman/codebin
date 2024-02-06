package searching

func BinarySearch(nums []int, target int) int {
    s := 0
    e := len(nums) - 1
    for s <= e {
        m := (s + e) / 2
        if nums[m] == target {
            return m
        } else if nums[m] < target {
            s++
        } else {
            e--
        }
    }
    return -1
}
