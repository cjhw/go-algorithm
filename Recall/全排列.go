package Recall

func permute(nums []int) [][]int {

	var (
		res  [][]int
		path []int
		st   []bool // state的缩写
	)
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))

	var dfs func(nums []int, cur int)

	dfs = func(nums []int, cur int) {
		if cur == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i := 0; i < len(nums); i++ {
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