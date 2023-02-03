package Greed

func partitionLabels(s string) []int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var res []int
	var mark [26]int
	size, left, right := len(s), 0, 0

	for i := 0; i < size; i++ {
		mark[s[i]-'a'] = i
	}

	for i := 0; i < size; i++ {
		right = max(right, mark[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}

	return res
}
