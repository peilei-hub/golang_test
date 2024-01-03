package main

import "fmt"

// https://leetcode.cn/problems/remove-k-digits

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)

	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, num[i])
	}

	stack = stack[:len(stack)-k]
	if len(stack) == 0 {
		return "0"
	}

	// 去除前面的'0'
	idx := 0
	for i, b := range stack {
		if b != '0' {
			idx = i
			break
		}
	}
	stack = stack[idx:]
	if stack[0] == '0' {
		return "0"
	}

	return string(stack)
}

func main() {
	kdigits := removeKdigits("10200", 1)
	fmt.Println(kdigits)
}
