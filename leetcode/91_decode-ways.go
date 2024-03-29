package main

import "fmt"

func numDecodings(s string) int {
	result := 0
	backtrace91(s, 0, &result)
	return result
}

func backtrace91(s string, start int, result *int) {
	if start >= len(s) {
		*result++
		return
	}

	u1 := s[start]

	if u1 == '0' {
		return
	}

	// 所有可能，一位 两位

	backtrace91(s, start+1, result) // 一位

	if start < len(s)-1 {
		u2 := s[start+1]
		if ((u1-'0')*10 + (u2 - '0')) <= 26 {
			backtrace91(s, start+2, result) // 两位
		}
	}
}

func main() {
	decodings := numDecodings("111111111111111111111111111111111111111111111")
	fmt.Println(decodings)
}
