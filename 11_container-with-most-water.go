package main

import "wangpeilei/leetcode/utils"

// https://leetcode.cn/problems/container-with-most-water/

// 若向内 移动短板 ，水槽的短板可能变大，因此下个水槽的面积 可能增大 。
// 若向内 移动长板 ，水槽的短板不变或变小，因此下个水槽的面积 一定变小 。

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	var max int
	for left < right {
		temp := utils.Min(height[left], height[right]) * (right - left)
		if temp > max {
			max = temp
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return max
}
