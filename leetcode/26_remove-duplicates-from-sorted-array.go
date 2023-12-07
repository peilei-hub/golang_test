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
	idx := 1 // 可以放置的位置下标

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx-1] { // 当前num跟idx的前一位不相同，可以将num放在idx位置
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}
