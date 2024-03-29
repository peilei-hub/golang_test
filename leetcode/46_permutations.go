package main

import "fmt"

// https://leetcode.cn/problems/permutations/

func permute(nums []int) [][]int {
	result := make([][]int, 0)

	tmp := make([]int, len(nums))
	used := make([]bool, len(nums))

	traceback46(&result, tmp, nums, used, 0)

	return result
}

func traceback46(result *[][]int, tmp []int, nums []int, used []bool, idx int) {
	if idx == len(nums) {
		t := make([]int, len(tmp))
		for i, v := range tmp {
			t[i] = v
		}
		*result = append(*result, t)
		return
	}

	for i, num := range nums {
		if used[i] {
			continue
		}

		used[i] = true
		tmp[idx] = num

		traceback46(result, tmp, nums, used, idx+1)

		used[i] = false
	}
}

func main() {
	res := permute([]int{1, 2, 3})

	fmt.Println(res)
}
