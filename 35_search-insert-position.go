package main

// https://leetcode.cn/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		}
	}

	return l
}
