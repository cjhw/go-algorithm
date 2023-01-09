package HashTable

import "sort"

func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	map_num := map[int]int{}

	for _, num := range nums {
		map_num[num]++
	}

	for key, _ := range map_num {
		ans = append(ans, key)
	}

	sort.Slice(ans, func(i, j int) bool {
		return map_num[ans[i]] > map_num[ans[j]]
	})

	return ans[:k]
}
