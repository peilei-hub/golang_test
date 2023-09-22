package main

import "sort"

// https://leetcode.cn/problems/merge-intervals/

func merge56(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	result := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}) // 排序，按照左区间的值升序，此时能合并的则必然相邻

	// 注意里面的 i<len(intervals)的判断
	for i := 0; i < len(intervals); {
		// 每次循环，找出能加到结果的一对
		leftMin, rightMax := intervals[i][0], intervals[i][1] // 初始范围
		i++                                                   // 从下一个开始
		for i < len(intervals) && intervals[i][0] <= rightMax {
			rightMax = max56(rightMax, intervals[i][1])
			i++
		}

		result = append(result, []int{leftMin, rightMax}) // 加到结果
	}

	return result
}

func max56(a, b int) int {
	if a > b {
		return a
	}

	return b
}
