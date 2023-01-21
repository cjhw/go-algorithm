package Recall

func letterCombinations(digits string) []string {
	var (
		m    []string
		path []byte
		res  []string
	)

	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res
	}

	var dfs func(dights string, start int)

	dfs = func(dights string, start int) {
		if len(path) == len(dights) {
			tmp := string(path)
			res = append(res, tmp)
			return
		}
		dight := int(dights[start] - '0')
		str := m[dight-2]
		for i := 0; i < len(str); i++ {
			path = append(path, str[i])
			dfs(dights, start+1)
			path = path[:len(path)-1]
		}
	}

	dfs(digits, 0)
	return res
}
