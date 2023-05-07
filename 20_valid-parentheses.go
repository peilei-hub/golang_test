package main

// https://leetcode.cn/problems/valid-parentheses/

// 思路，左括号入栈

func isValid(s string) bool {
	stack := make([]byte, 0)

	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, b := range s {
		if b == '(' || b == '{' || b == '[' { // 左括号就入栈
			stack = append(stack, byte(b))
			continue
		}

		// b是右括号

		// 此时栈里没东西，则不合法
		if len(stack) == 0 {
			return false
		}

		// 判断栈里的头元素左括号是否跟当前的右括号匹配
		if m[stack[len(stack)-1]] == byte(b) {
			stack = stack[:len(stack)-1] // 匹配就抵消掉，将左括号从栈里去掉
		} else {
			return false
		}
	}

	// 最终栈里为空，表示合法
	return len(stack) == 0
}
