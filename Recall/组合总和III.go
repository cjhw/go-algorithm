package Recall

func combinationSum3(k int, n int) [][]int {

	var (
		path []int
		res  [][]int
	)

	res, path = make([][]int, 0), make([]int, 0)

	var dfs func(k int, n int, start int, sum int)

	dfs = func(k, n, start, sum int) {
		if len(path) == k && sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)

			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n || 9-i+i < k-len(path) {
				break
			}
			path = append(path, i)
			dfs(k, n, i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	dfs(k, n, 1, 0)

	return res
}
