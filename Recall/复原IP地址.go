package Recall

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var (
		path []string
		res  []string
	)

	path, res = make([]string, 0, len(s)), make([]string, 0)

	var dfs func(s string, start int)

	dfs = func(s string, start int) {
		if len(path) == 4 {
			if start == len(s) {
				str := strings.Join(path, ".")
				res = append(res, str)
			}
			return
		}
		for i := start; i < len(s); i++ {
			if i != start && s[start] == '0' { // 含有前导 0，无效
				break
			}
			str := s[start : i+1]
			num, _ := strconv.Atoi(str)
			if num <= 255 && num >= 0 {
				path = append(path, str)
				dfs(s, i+1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}

	}

	dfs(s, 0)
	return res
}