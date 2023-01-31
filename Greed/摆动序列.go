package Greed

func wiggleMaxLength(nums []int) int {

	n := len(nums)

	if n < 2 {
		return n
	}
	res := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		res = 2
	}

	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if prevDiff >= 0 && diff < 0 || prevDiff <= 0 && diff > 0 {
			res++
			prevDiff = diff
		}
	}

	return res
}