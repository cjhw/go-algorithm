package string

// func replaceSpace(s string) string {

// 	b := []byte(s)
// 	result := make([]byte, 0)

// 	for _, v := range b {
// 		if v == ' ' {
// 			result = append(result, []byte("%20")...)
// 		} else {
// 			result = append(result, v)
// 		}
// 	}

// 	return string(result)
// }

func replaceSpace(s string) string {

	b := []byte(s)
	oldLength := len(s)
	spaceNum := 0

	for i := 0; i < oldLength; i++ {
		if b[i] == ' ' {
			spaceNum += 1
		}
	}

	tmp := make([]byte, spaceNum*2)
	b = append(b, tmp...)
	i := oldLength - 1
	j := len(b) - 1

	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}
