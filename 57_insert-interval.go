package main

// https://leetcode.cn/problems/insert-interval/

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)

	left, right := newInterval[0], newInterval[1]

	added := false
	for _, interval := range intervals {
		if added {
			result = append(result, interval)
			continue
		}

		curLeft, curRight := interval[0], interval[1]

		if curRight < left {
			result = append(result, []int{curLeft, curRight})
		} else if curLeft > right {
			result = append(result, []int{left, right})
			added = true
			result = append(result, []int{curLeft, curRight})
		} else {
			left = min57(left, curLeft)
			right = max57(right, curRight)
		}
	}

	if !added {
		result = append(result, []int{left, right})
	}

	return result
}

func min57(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max57(a, b int) int {
	if a > b {
		return a
	}
	return b
}
