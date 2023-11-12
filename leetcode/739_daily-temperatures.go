package main

// https://leetcode.cn/problems/daily-temperatures/description/

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	stack := make([]int, 0)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i) // 下标入栈，用下标判断
	}
	return result
}
