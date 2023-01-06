package HashTable

func twoSum(nums []int, target int) []int {

	help := make(map[int]int, 0)
	for index, val := range nums {
		if pre, ok := help[target-nums[index]]; ok {
			return []int{pre, index}
		} else {
			help[val] = index
		}
	}

	return []int{}
}
