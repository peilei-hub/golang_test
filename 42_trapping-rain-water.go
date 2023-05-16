package main

// https://leetcode.cn/problems/trapping-rain-water/

func trap(height []int) int {
	// 维护两个数组，分别为当前下标的左边最大值和当前下标的右边最大值
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	for i := 1; i < len(height); i++ {
		leftMax[i] = max42(leftMax[i-1], height[i-1]) // 当前下标的左边最大值为："左边的左边的最大值"和"左边的高度"的最大值
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max42(rightMax[i+1], height[i+1])
	}

	total := 0
	for i := 1; i < len(height)-1; i++ {
		tmp := min42(leftMax[i], rightMax[i]) - height[i]
		if tmp > 0 {
			total += tmp
		}
	}
	return total
}

func min42(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max42(a, b int) int {
	if a > b {
		return a
	}
	return b
}
