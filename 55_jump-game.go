package main

// https://leetcode.cn/problems/jump-game/

func canJump(nums []int) bool {
	nextBorder := 0
	curBorder := 0
	for i := 0; i < len(nums)-1; i++ {
		nextBorder = max55(nextBorder, nums[i]+i)

		if i == curBorder {
			if curBorder == nextBorder {
				return false
			}
			curBorder = nextBorder
		}
	}

	return true
}

func max55(a, b int) int {
	if a > b {
		return a
	}
	return b
}
