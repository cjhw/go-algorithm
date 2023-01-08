package Array

func minSubArrayLen(target int, nums []int) int {
	i := 0
	length := len(nums)
	res := length + 1
	sum := 0

	for j := 0; j < length; j++ {
		sum += nums[j]
		for sum >= target {
			subLen := j - i + 1
			if res > subLen {
				res = subLen
			}
			sum -= nums[i]
			i++
		}
	}

	if res >= length+1 {
		return 0
	} else {
		return res
	}
}
