package main

// https://leetcode.cn/problems/longest-valid-parentheses/

// 将左括号对应的下标入栈，维护一个validArray []bool 存储对应的下边括号是否有效

func longestValidParentheses(s string) int {
	stack := make([]int, 0)

	validArray := make([]bool, len(s))
	for i := range validArray { // 初始化
		validArray[i] = true
	}

	for i, b := range s {
		if b == '(' {
			stack = append(stack, i) // 左括号，将下标入栈
		} else { // b = ')'
			if len(stack) == 0 {
				validArray[i] = false // 如果栈为空，表示里面没有左括号，对应的位置为无效括号
				continue
			}

			stack = stack[:len(stack)-1] // 将一个左括号出栈，消耗一个左括号
		}
	}

	for len(stack) > 0 { // 栈里多的左括号，对应的位置都为无效括号
		i := stack[len(stack)-1]
		validArray[i] = false
		stack = stack[:len(stack)-1]
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
