package main

import "fmt"

// https://leetcode.cn/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	greater := nextGreater(nums2)

	numsMap := make(map[int]int, len(nums2))
	for i, v := range nums2 {
		numsMap[v] = greater[i]
	}

	for i, v := range nums1 {
		result[i] = numsMap[v]
	}
	return result
}

func nextGreater(nums []int) []int {
	result := make([]int, len(nums))

	stack := make([]int, 0)

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] { // 判定个子高矮
			stack = stack[:len(stack)-1] // 矮个子不要，被挡着了
		}

		if len(stack) == 0 { // 栈里没有比nums[i]更大的
			result[i] = -1
		} else { // 有
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	return result
}

func main() {
	nums := []int{2, 1, 2, 4, 3}

	greater := nextGreater(nums)

	fmt.Println(greater)
}
