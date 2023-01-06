package HashTable

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]string, 0)
	res := make([]int, 0)

	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = "1"
		}

	}

	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}

	return res
}
