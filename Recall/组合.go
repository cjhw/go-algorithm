package Recall

func combine(n int, k int) [][]int {
	var (
		path []int
		res  [][]int
	)

	path, res = make([]int, 0, k), make([][]int, 0)

	var dfs func(n int, k int, start int)

	dfs = func(n int, k int, start int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i <= n; i++ {
			if n-i+1 < k-len(path) { // 剪枝
				break
			}
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	dfs(n, k, 1)
	return res
}
