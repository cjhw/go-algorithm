package Recall

func findSubsequences(nums []int) [][]int {
	var (
		path []int
		res  [][]int
	)

	res, path = make([][]int, 0), make([]int, 0, len(nums))

	var dfs func(nums []int, start int)

	dfs = func(nums []int, start int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		used := make(map[int]bool, len(nums))
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[nums[i]] = true
				dfs(nums, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(nums, 0)

	return res
}