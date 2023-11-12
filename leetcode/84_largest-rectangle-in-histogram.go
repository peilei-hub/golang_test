package main

import "fmt"

// https://leetcode.cn/problems/largest-rectangle-in-histogram/description/

// 暴力遍历
func largestRectangleArea1(heights []int) int {
	max := 0
	curHeight := 0
	for i := 0; i < len(heights); i++ {
		curHeight = heights[i]

		left, right := i-1, i+1

		for left >= 0 && heights[left] >= curHeight {
			left--
		}

		for right <= len(heights)-1 && heights[right] >= curHeight {
			right++
		}

		curArea := curHeight * (right - left - 1)
		max = max84(curArea, max)
	}
	return max
}

func max84(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea2(heights []int) int {
	preSmallerNums := preSmaller(heights)
	nextSmallerNums := nextSmaller(heights)

	max := 0
	for i := 0; i < len(heights); i++ {
		pre, next := preSmallerNums[i], nextSmallerNums[i]

		curArea := heights[i] * (next - pre - 1)
		max = max84(max, curArea)
	}
	return max
}

func nextSmaller(nums []int) []int {
	result := make([]int, len(nums))

	stack := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = len(nums) // 下标为len(nums)
		} else {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}

func preSmaller(nums []int) []int {
	result := make([]int, len(nums))

	stack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = -1 // 下标为-1
		} else {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return result
}

func main() {
	heights := []int{5, 4, 1, 2}
	area := largestRectangleArea2(heights)
	fmt.Println(area)

	smaller1 := nextSmaller(heights)
	fmt.Println(smaller1)

	smaller2 := preSmaller(heights)
	fmt.Println(smaller2)
}
