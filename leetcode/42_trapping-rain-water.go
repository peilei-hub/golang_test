package main

import "wangpeilei/leetcode/utils"

// https://leetcode.cn/problems/trapping-rain-water/

func trap(height []int) int {
	// 维护两个数组，分别为当前下标的左边最大值和当前下标的右边最大值
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	for i := 1; i < len(height); i++ {
		leftMax[i] = utils.Max(leftMax[i-1], height[i-1]) // 当前下标的左边最大值为："左边的左边的最大值"和"左边的高度"两者的最大值
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = utils.Max(rightMax[i+1], height[i+1]) // 当前下标的右边最大值为："右边的右边的最大值"和"右边的高度"两者的最大值
	}

	total := 0
	for i := 1; i < len(height)-1; i++ { // 去掉头尾
		tmp := utils.Min(leftMax[i], rightMax[i]) - height[i]
		if tmp > 0 {
			total += tmp
		}
	}
	return total
}
