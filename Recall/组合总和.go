package Recall

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var (
		res  [][]int
		path []int
	)

	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates)

	var dfs func(candidates []int, start int, target int)

	dfs = func(candidates []int, start, target int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				// 剪枝
				break
			}

			path = append(path, candidates[i])
			dfs(candidates, i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}

	dfs(candidates, 0, target)
	return res
}
