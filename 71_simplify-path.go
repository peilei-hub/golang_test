package main

import "strings"

// https://leetcode.cn/problems/simplify-path/

// 把当前目录压到栈，遇到".."出栈，最后返回栈中元素

// 用path根据"/"分割出现4种情况
// 1. 空字符串 2. 一个点"." 3. 两点".."  4. 目录名

func simplifyPath(path string) string {
	stack := make([]string, 0) // 栈里保存的目录名

	split := strings.Split(path, "/")
	for _, s := range split {
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 返回上一级，要将目录出栈
			}
		} else if s != "" && s != "." {
			stack = append(stack, s) // 目录名
		}
	}

	return "/" + strings.Join(stack, "/")
}
