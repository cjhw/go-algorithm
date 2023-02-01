package Greed

func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1

	var max func(a, b int) int

	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= n {
			return true
		}
	}

	return false

}
