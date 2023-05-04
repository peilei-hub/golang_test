package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	var idx int

	for _, num := range nums {
		if nums[idx] != num {
			idx++
			nums[idx] = num
		}
	}

	return idx + 1
}
