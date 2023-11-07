package main

// https://leetcode.cn/problems/house-robber-ii/description/

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max213(nums[0], nums[1])
	}

	nums1 := nums[:len(nums)-1]
	nums2 := nums[1:]
	return max213(rob1(nums1), rob1(nums2))
}

func rob1(nums []int) int {
	if l := len(nums); l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		return max213(nums[0], nums[1])
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = max213(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}
