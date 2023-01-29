package Recall

func subsets(nums []int) [][]int {

	var (
		path []int
		res  [][]int
	)

	res, path = make([][]int, 0), make([]int, 0, len(nums))
	var dfs func(nums []int, start int)

	dfs = func(nums []int, start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return res
}