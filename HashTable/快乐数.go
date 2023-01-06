package HashTable

func isHappy(n int) bool {
	m := make(map[int]bool, 0)
	for n != 1 && !m[n] {
		n, m[n] = getNum(n), true
	}

	return n == 1
}

func getNum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}

	return sum
}
