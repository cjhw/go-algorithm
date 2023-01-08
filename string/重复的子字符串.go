package String

import "fmt"

func RepeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	next := make([]int, n)

	j := 0
	next[0] = j

	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}

		next[i] = j
	}

	fmt.Println((next[n-1]))

	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}

	return false
}
