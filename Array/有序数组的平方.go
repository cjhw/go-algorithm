package array

func sortedSquares(nums []int) []int {
	length := len(nums)
	left, right, k := 0, length-1, length-1
	res := make([]int, length)
	for left <= right {
		l, r := nums[left]*nums[left], nums[right]*nums[right]
		if l > r {
			res[k] = l
			left++
		} else {
			res[k] = r
			right--
		}
		k--
	}

	return res
}
