package main

func search(nums []int, target int) int {
	low := 0
	high := len(nums)
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
