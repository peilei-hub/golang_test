package main

import "fmt"

// https://leetcode.cn/problems/longest-palindromic-substring/

// 动态规划 dp[i][j]为 i-j 为回文
// [i,j]里面有 i,i+1,j-1,j。i,j又依赖i+1,j-1。编历时先限定j，然后i从j到0
func longestPalindrome1(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	var maxLength, start int

	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = true // i=j 一定为回文
			} else {
				if s[i] == s[j] {
					if i+1 >= j-1 { //
						dp[i][j] = true
					} else {
						dp[i][j] = dp[i+1][j-1]
					}
					// 跟上面判断一样
					//if j-i < 3 { //  长度小于 3 一定为回文
					//	dp[i][j] = true
					//} else {
					//	dp[i][j] = dp[i+1][j-1]
					//}
				} else {
					dp[i][j] = false
				}
			}

			if dp[i][j] { // 更新 maxLength
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
		str1 := expandPalindrome(s, i, i)   //  以 i 为中心点
		str2 := expandPalindrome(s, i, i+1) //  以 i,i+1 为中心点

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
	palindrome2 := longestPalindrome1("babad")
	fmt.Println(palindrome2)
}
