package main

import "sort"

// https://leetcode.cn/problems/merge-intervals/

func merge56(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	result := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}) // 排序，按照左区间的值升序，此时能合并的则必然相邻

	lMinVal, rMaxVal := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {

		// 如果左区间小于 rMaxVal，则一定能合并。此时更新rMaxVal
		for i < len(intervals) && intervals[i][0] <= rMaxVal {
			rMaxVal = max56(rMaxVal, intervals[i][1])
			i++
		}

		result = append(result, []int{lMinVal, rMaxVal}) // 加到结果

		if i < len(intervals) {
			lMinVal, rMaxVal = intervals[i][0], intervals[i][1] // 新的一轮
		}
	}

	if rMaxVal > result[len(result)-1][1] { // 还有剩下没有加进来的
		result = append(result, []int{lMinVal, rMaxVal})
	}

	return result
}

func max56(a, b int) int {
	if a > b {
		return a
	}

	return b
}
