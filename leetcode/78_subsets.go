package main

import "fmt"

// https://leetcode.cn/problems/subsets/

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	backtrack78(&result, nums, []int{}, 0)
	return result
}

func backtrack78(result *[][]int, nums []int, tmp []int, start int) {
	*result = append(*result, append([]int{}, tmp...))

	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])

		backtrack78(result, nums, tmp, i+1)

		tmp = tmp[:len(tmp)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}

	result := subsets(nums)

	fmt.Println(result)
}
