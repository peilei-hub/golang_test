package main

// https://leetcode.cn/problems/remove-element/

func removeElement(nums []int, val int) int {
	var idx int
	for _, num := range nums {
		if num != val {
			nums[idx] = num
			idx++
		}
	}

	return idx
}
