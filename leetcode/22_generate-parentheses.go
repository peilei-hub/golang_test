package main

import "fmt"

// https://leetcode.cn/problems/generate-parentheses/

// 左右括号计数

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	result := make([]string, 0)

	tmp := make([]byte, n*2)
	leftNum, rightNum := n, n

	backtrace22(&result, tmp, leftNum, rightNum, 2*n, 0)

	return result
}

func backtrace22(result *[]string, tmp []byte, leftNum, rightNum, total, idx int) {
	if idx == total { // 到截止条件，返回
		*result = append(*result, string(tmp))
		return
	}

	if leftNum < rightNum { // 如果左括号数量少于右括号，则可以添加左括号，也可以添加右括号，需要遍历两个可能
		if leftNum > 0 {
			tmp[idx] = '('
			backtrace22(result, tmp, leftNum-1, rightNum, total, idx+1)
		}
		tmp[idx] = ')'
		backtrace22(result, tmp, leftNum, rightNum-1, total, idx+1)
	} else if leftNum == rightNum { // 如果左括号数量等于右括号，只能添加左括号
		tmp[idx] = '('
		backtrace22(result, tmp, leftNum-1, rightNum, total, idx+1)
	}
	// 保证不存在左括号数量多于右括号的情况
}

func main() {
	parenthesis := generateParenthesis(3)
	fmt.Println(parenthesis)
}
