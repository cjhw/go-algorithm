package DoublePointer

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}

		if i > 0 && n1 == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			n2 := nums[l]
			n3 := nums[r]

			total := n1 + n2 + n3

			if total == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && n2 == nums[l] {
					l++
				}
				for l < r && n3 == nums[r] {
					r--
				}
			} else if total < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
