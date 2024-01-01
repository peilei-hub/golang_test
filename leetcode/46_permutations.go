package main

import "fmt"

// https://leetcode.cn/problems/permutations/

func permute(nums []int) [][]int {
	result := make([][]int, 0)

	tmp := make([]int, len(nums))
	used := make([]bool, len(nums))

	backtrack46(&result, tmp, nums, used, 0)

	return result
}

func backtrack46(result *[][]int, tmp []int, nums []int, used []bool, idx int) {
	if idx == len(nums) {
		*result = append(*result, append([]int{}, tmp...))
		return
	}

	for i, num := range nums {
		if used[i] {
			continue
		}

		used[i] = true
		tmp[idx] = num

		backtrack46(result, tmp, nums, used, idx+1)

		used[i] = false
	}
}

func main() {
	res := permute([]int{1, 2, 3})

	fmt.Println(res)
}
