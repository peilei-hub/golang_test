package main

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {

	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {

			l1 := m
			for l1 > 0 && nums[l1] == nums[l1-1] {
				l1--
			}

			r1 := m
			for r1 < len(nums)-1 && nums[r1] == nums[r1+1] {
				r1++
			}

			return []int{l1, r1}

		} else if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		}
	}

	return []int{-1, -1}

}
