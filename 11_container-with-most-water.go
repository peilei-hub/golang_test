package main

// https://leetcode.cn/problems/container-with-most-water/

// 若向内 移动短板 ，水槽的短板可能变大，因此下个水槽的面积 可能增大 。
// 若向内 移动长板 ，水槽的短板不变或变小，因此下个水槽的面积 一定变小 。

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		max := minInt(height[l], height[r]) * (l - r)
		if max > res {
			res = max
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return res
}

func minInt(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
