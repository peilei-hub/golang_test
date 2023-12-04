package main

// https://leetcode.cn/problems/house-robber/

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max198(nums[0], nums[1])
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		// 前天偷+今天，或者昨天偷
		dp[i] = max198(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[length]
}

func max198(a, b int) int {
	if a > b {
		return a
	}
	return b
}
