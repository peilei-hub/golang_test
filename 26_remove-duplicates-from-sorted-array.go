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

func removeDuplicates1(nums []int) int {
	idx := 1

	for i := 1; i < len(nums); i++ {
		if nums[idx-1] != nums[i] {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}
