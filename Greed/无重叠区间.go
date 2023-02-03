package Greed

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			res++
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
		}
	}

	return res
}
