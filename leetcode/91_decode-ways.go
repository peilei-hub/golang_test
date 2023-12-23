package main

import "fmt"

// https://leetcode.cn/problems/decode-ways/

func numDecodings(s string) int {
	result := 0
	backtrack91(s, 0, &result)
	return result
}

func backtrack91(s string, start int, result *int) {
	if start >= len(s) {
		*result++
		return
	}

	u1 := s[start]

	if u1 == '0' {
		return
	}

	// 所有可能，一位 两位

	backtrack91(s, start+1, result) // 一位

	if start < len(s)-1 {
		u2 := s[start+1]
		if ((u1-'0')*10 + (u2 - '0')) <= 26 {
			backtrack91(s, start+2, result) // 两位
		}
	}
}

func numDecodings2(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		temp := 0
		if s[i-1] != '0' {
			temp += dp[i-1]
		}

		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			temp += dp[i-2]
		}

		dp[i] = temp
	}

	return dp[n]
}

func main() {
	//decodings := numDecodings2("111111111111111111111111111111111111111111111")
	decodings := numDecodings2("12")
	fmt.Println(decodings)
}
