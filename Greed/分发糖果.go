package Greed

func candy(ratings []int) int {
	need := make([]int, len(ratings))

	var findMax func(a, b int) int
	findMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(need); i++ {
		need[i] = 1
	}
	// 1.先从左到右，当右边的大于左边的就加1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			need[i+1] = need[i] + 1
		}
	}
	// 2.再从右到左，当左边的大于右边的就右边加1，但要花费糖果最少，所以需要做下判断
	for i := len(ratings) - 1; i > 0; i-- {

		if ratings[i-1] > ratings[i] {
			need[i-1] = findMax(need[i]+1, need[i-1])
		}

	}

	sum := 0
	for i := 0; i < len(need); i++ {
		sum += need[i]
	}

	return sum

}