package main

// https://leetcode.cn/problems/jump-game/

func canJump(nums []int) bool {
	nextBorder := 0
	curBorder := 0
	for i := 0; i < len(nums)-1; i++ {
		nextBorder = max55(nextBorder, nums[i]+i) // 在当前border内去找下次的border

		if i == curBorder { // 到达当前border
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
