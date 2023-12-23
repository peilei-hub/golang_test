package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	tmp := make([]int, len(nums))
	used := make([]bool, len(nums))

	traceback47(&result, tmp, used, nums, 0)
	return result
}

func traceback47(result *[][]int, tmp []int, used []bool, nums []int, idx int) {
	if idx == len(tmp) {
		*result = append(*result, append([]int{}, tmp...))
		return
	}

	for i, num := range nums {
		if used[i] || (!used[i-1] && i > 0 && nums[i] == nums[i-1]) { // 如果前一个数没被选过，那也不能选当前的数
			continue
		}

		tmp[idx] = num
		used[i] = true

		traceback47(result, tmp, used, nums, idx+1)

		used[i] = false
	}
}

func main() {
	unique := permuteUnique([]int{1, 1, 2})
	fmt.Println(unique)
}
