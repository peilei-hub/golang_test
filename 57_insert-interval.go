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

		if curRight < left { // 如果curRight < left, 还没到newInterval的范围, 那直接加进去
			result = append(result, []int{curLeft, curRight})
		} else if curLeft > right { // 如果curLeft > right，即刚过newInterval的范围
			result = append(result, []int{left, right}) // 将left right加进去，并设置加入flag。
			added = true
			result = append(result, []int{curLeft, curRight})
		} else { // 如果 curLeft <= right && curRight >= left
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
