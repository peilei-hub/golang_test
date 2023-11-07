package main

// https://leetcode.cn/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	res := 0

	i := len(s) - 1
	for s[i] != ' ' && i >= 0 {
		i--
	}
	if i < 0 {
		return 0
	}

	for j := i; j >= 0; j-- {

		if s[j] == ' ' {
			return res
		}
		res++
	}
	return res
}
