package main

// https://leetcode.cn/problems/longest-valid-parentheses/

// 将左括号对应的下标入栈，维护一个validArray []bool 存储对应的下边括号是否有效

func longestValidParentheses(s string) int {
	validArray := make([]bool, len(s))
	for i := range validArray {
		validArray[i] = true // 都初始化为true
	}

	stack := make([]int, 0)
	for i, b := range s {
		if b == '(' {
			stack = append(stack, i) // 将左括号下标入栈
		} else { // b == ')'
			if len(stack) == 0 { // 没有左括号能匹配
				validArray[i] = false
			} else {
				stack = stack[:len(stack)-1] // 出栈一个
			}
		}
	}

	for _, n := range stack {
		validArray[n] = false // 还在栈里的是没有匹配的左括号
	}

	max := 0
	tmp := 0
	for _, b := range validArray {
		if b {
			tmp++
			max = maxInt(max, tmp)
		} else {
			tmp = 0
		}
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
