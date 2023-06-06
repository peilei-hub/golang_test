package main

import (
	"fmt"
	"math/big"
)

// https://leetcode.cn/problems/unique-paths/

func uniquePathsV1(m int, n int) int {
	// dp[i][j] 是到达i, j的最多路径
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// 排列组合 C
// m−1
// m+n−2
func uniquePathsV2(m int, n int) int {
	a := m - 1
	if m > n {
		a = n - 1
	}
	b := m + n - 2

	total := big.NewInt(int64(1))

	for i := 0; i < a; i++ {
		total.Mul(total, big.NewInt(int64(b-i)))
	}

	for i := 1; i <= a; i++ {
		total.Div(total, big.NewInt(int64(i)))
	}

	return int(total.Int64())
}

func main() {
	v2 := uniquePathsV2(3, 7)
	fmt.Println(v2)
}
