package Greed

func jump(nums []int) int {

	var max func(a, b int) int

	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(nums)
	if n == 1 {
		return n
	}

	cur := 0
	next := 0
	step := 0

	for i := 0; i < n-1; i++ {
		next := max(nums[i]+i, next)
		if i == cur {
			cur = next
			step++
		}
	}

	return step
}
