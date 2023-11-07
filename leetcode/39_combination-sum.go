package main

import "fmt"

// https://leetcode.cn/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	traceback39(candidates, target, &result, nil, 0, 0)

	return result
}

func traceback39(candidates []int, target int, result *[][]int, tmp []int, sum int, begin int) {
	if sum == target { // 如果找到
		array := make([]int, len(tmp))
		for i := range tmp {
			array[i] = tmp[i]
		}
		*result = append(*result, array)
		return
	}
	if sum > target { // 如果大于，则后面不可能找到，直接返回
		return
	}

	// sum < target
	for i := begin; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		traceback39(candidates, target, result, tmp, sum+candidates[i], i)

		tmp = tmp[:len(tmp)-1]
	}
}

func main() {
	sum := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(sum)
}
