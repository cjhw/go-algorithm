package Greed

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	l := len(ss)
	for i := l - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < l; j++ {
				ss[j] = '9'
			}
		}
	}

	res, _ := strconv.Atoi(string(ss))

	return res
}
