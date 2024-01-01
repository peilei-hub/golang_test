package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/permutation-sequence/

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	used := make([]bool, n)

	result := ""
	tmp := make([]int, n)

	total := 0
	backtrack60(&result, tmp, nums, used, 0, k, &total)

	fmt.Println(result)

	return result
}

func backtrack60(result *string, tmp []int, nums []int, used []bool, idx int, k int, total *int) {
	if *total >= k {
		return
	}
	if idx == len(nums) {
		*total++

		if *total == k {
			str := ""
			for _, i := range tmp {
				str += strconv.Itoa(i)
			}
			*result = str
		}
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			tmp[idx] = nums[i]

			backtrack60(result, tmp, nums, used, idx+1, k, total)

			used[i] = false
		}
	}
}

func main() {
	permutation := getPermutation(4, 9)
	fmt.Println(permutation)
}
