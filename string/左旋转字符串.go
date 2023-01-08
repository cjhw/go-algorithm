package string

func reverseLeftWords(s string, n int) string {
	b := []byte(s)

	reverse1(b, 0, n-1)
	reverse1(b, n, len(b)-1)
	reverse1(b, 0, len(b)-1)

	return string(b)
}
