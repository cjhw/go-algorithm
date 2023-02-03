package Greed

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := make([][]int, 0, len(intervals))
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(intervals[i][1], right)
		}
	}

	res = append(res, []int{left, right})

	return res
}
