package HashTable

func IsAnagram(s string, t string) bool {
	temp := [26]int{}

	for _, r := range s {
		temp[r-rune('a')]++
	}

	for _, r := range t {
		temp[r-rune('a')]--
	}

	return temp == [26]int{}
}
