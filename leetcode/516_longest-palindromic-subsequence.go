package main

import "fmt"

// 动态规划 dp[i][j]为从i到j最长回文子序列
// [i,j]里面有 i,i+1,j-1,j。i,j又依赖i+1,j-1。编历时先限定j，然后i从j到0
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = 1
			} else {
				if s[i] == s[j] {
					if i+1 == j-1 {
						dp[i][j] = 2 + dp[i+1][j-1] // 此时dp[i+1][j-1]为1
					} else if i+1 > j-1 {
						dp[i][j] = 2 + dp[i+1][j-1] // 此时dp[i+1][j-1]为0
					} else {
						dp[i][j] = 2 + dp[i+1][j-1]
					}
					// 可以简化为 dp[i][j] = dp[i+1][j-1] + 2
				} else {
					max := dp[i][j-1]
					if dp[i+1][j] > max {
						max = dp[i+1][j]
					}
					dp[i][j] = max
				}
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	str := "bbbab"
	subseq := longestPalindromeSubseq(str)
	fmt.Println(subseq)
}
