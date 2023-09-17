package main

import "fmt"

// https://leetcode.cn/problems/subsets/

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	tmp := make([]int, 0)

	traceback78(&result, nums, &tmp, 0)
	return result
}

func traceback78(result *[][]int, nums []int, tmp *[]int, start int) {
	cur := make([]int, len(*tmp))
	for i, v := range *tmp {
		cur[i] = v
	}
	*result = append(*result, cur)

	for i := start; i < len(nums); i++ {
		*tmp = append(*tmp, nums[i])

		traceback78(result, nums, tmp, i+1)

		*tmp = (*tmp)[:len(*(tmp))-1]
	}
}

func main() {
	nums := []int{1, 2, 3}

	result := subsets(nums)

	fmt.Println(result)
}
