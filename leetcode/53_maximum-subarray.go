package main

// https://leetcode.cn/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	var max, sum = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		max = max53(max, sum)
	}

	return max
}

func max53(a, b int) int {
	if a > b {
		return a
	}
	return b
}
