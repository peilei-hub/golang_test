package main

// https://leetcode.cn/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	max := nums[0]

	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
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
