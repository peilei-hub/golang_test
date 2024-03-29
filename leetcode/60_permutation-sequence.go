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
	traceback60(&result, tmp, nums, used, 0, k, &total)

	fmt.Println(result)

	return result
}

func traceback60(result *string, tmp []int, nums []int, used []bool, step int, k int, total *int) {
	if *total >= k {
		return
	}
	if step == len(nums) {
		*total++
		str := ""
		for _, i := range tmp {
			str += strconv.Itoa(i)
		}

		if *total == k {
			*result = str
		}
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			tmp[step] = nums[i]

			traceback60(result, tmp, nums, used, step+1, k, total)

			used[i] = false
		}
	}
}

func main() {
	permutation := getPermutation(4, 9)
	fmt.Println(permutation)
}
