package string

func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)

	for i := 0; i < length; i += 2 * k {
		if i+k >= length {
			reverse(ss[i:length])
		} else {
			reverse(ss[i : i+k])
		}
	}

	return string(ss)
}

func reverse(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
