package Recall

import "sort"

func permuteUnique(nums []int) [][]int {
	var (
		res  [][]int
		path []int
		st   []bool
	)

	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))

	sort.Ints(nums)

	var dfs func(nums []int, cur int)

	dfs = func(nums []int, cur int) {
		if cur == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i := 0; i < len(nums); i++ {
			if i != 0 && nums[i] == nums[i-1] && !st[i-1] {
				continue
			}

			if !st[i] {
				path = append(path, nums[i])
				st[i] = true
				dfs(nums, cur+1)
				st[i] = false
				path = path[:len(path)-1]
			}
		}
	}

	dfs(nums, 0)

	return res
}