package string

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			if j == len(needle)-1 {
				return i - j
			} else {
				j++
			}
		}
	}

	return -1
}

func getNext(next []int, s string) []int {
	j := 0
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}

	return next
}
