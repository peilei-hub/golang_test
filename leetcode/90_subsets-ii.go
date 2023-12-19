package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/subsets-ii/description/

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums)

	backtrack90(nums, []int{}, &result, 0)
	return result
}

func backtrack90(nums []int, temp []int, result *[][]int, start int) {
	t := make([]int, len(temp))
	copy(t, temp)
	*result = append(*result, t)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		temp = append(temp, nums[i])

		backtrack90(nums, temp, result, i+1)

		temp = temp[:len(temp)-1]
	}
}

func main() {
	nums := []int{4, 4, 4, 1, 4}

	dup := subsetsWithDup(nums)

	fmt.Println(dup)
}
