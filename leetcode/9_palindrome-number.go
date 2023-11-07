package main

// https://leetcode.cn/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	y := x
	var res int
	for y != 0 {
		cur := y % 10
		y /= 10

		res = res*10 + cur
	}
	return res == x
}
