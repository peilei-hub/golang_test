package main

import "fmt"

// https://leetcode.cn/problems/subsets/

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	traceback78(&result, nums, []int{}, 0)
	return result
}

func traceback78(result *[][]int, nums []int, tmp []int, start int) {
	cur := make([]int, len(tmp))
	copy(cur, tmp)
	*result = append(*result, cur)

	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])

		traceback78(result, nums, tmp, i+1)

		tmp = tmp[:len(tmp)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}

	result := subsets(nums)

	fmt.Println(result)
}
