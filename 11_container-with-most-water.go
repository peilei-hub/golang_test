package main

// https://leetcode.cn/problems/container-with-most-water/

// 若向内 移动短板 ，水槽的短板可能变大，因此下个水槽的面积 可能增大 。
// 若向内 移动长板 ，水槽的短板不变或变小，因此下个水槽的面积 一定变小 。

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	var max int
	for left < right {
		temp := minInt(height[left], height[right]) * (right - left)
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

func minInt(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
