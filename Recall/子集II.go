package Recall

import "sort"

func subsetsWithDup(nums []int) [][]int {

	var (
		path []int
		res  [][]int
	)

	path, res = make([]int, 0, len(nums)), make([][]int, 0)
	sort.Ints(nums)

	var dfs func(nums []int, start int)

	dfs = func(nums []int, start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return res
}