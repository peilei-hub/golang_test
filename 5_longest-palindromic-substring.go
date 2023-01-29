package main

import "fmt"

// https://leetcode.cn/problems/longest-palindromic-substring/

func longestPalindrome1(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	var maxLength, start int

	for j := range s {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
			} else {
				if s[i] == s[j] {
					if j-i < 3 {
						dp[i][j] = true
					} else {
						dp[i][j] = dp[i+1][j-1]
					}
				} else {
					dp[i][j] = false
				}
			}

			if dp[i][j] {
				if j-i+1 > maxLength {
					maxLength = j - i + 1
					start = i
				}
			}
		}
	}

	return s[start : start+maxLength]
}

func longestPalindrome2(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var result string
	for i := 0; i < len(s); i++ {
		str1 := expandPalindrome(s, i, i)
		str2 := expandPalindrome(s, i, i+1)

		if len(str1) > len(str2) {
			if len(str1) > len(result) {
				result = str1
			}
		} else {
			if len(str2) > len(result) {
				result = str2
			}
		}
	}

	return result
}

func expandPalindrome(s string, i, j int) string {
	if i == len(s)-1 { // i是最后一位，特殊处理
		return s[i:]
	}

	for i >= 0 && j < len(s) {
		if s[i] == s[j] { // 相等就左右各扩大一格范围
			i--
			j++
		} else {
			break
		}
	}

	return s[i+1 : j] // 上面扩大了范围，需要去掉
}

func main() {
	palindrome2 := longestPalindrome2("babad")
	fmt.Println(palindrome2)
}
