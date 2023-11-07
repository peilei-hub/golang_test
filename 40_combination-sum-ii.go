package main

import "sort"

// https://leetcode.cn/problems/combination-sum-ii/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := make([][]int, 0)

	traceback40(candidates, &result, nil, target, 0, 0)
	return result
}

func traceback40(candidates []int, result *[][]int, tmp []int, target int, sum int, begin int) {
	if sum == target {
		t := make([]int, len(tmp))
		for i, v := range tmp {
			t[i] = v
		}

		*result = append(*result, t)
		return
	}
	if sum > target {
		return
	}

	// sum < target
	for i := begin; i < len(candidates); i++ {
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}

		tmp = append(tmp, candidates[i])
		traceback40(candidates, result, tmp, target, sum+candidates[i], i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
