package main

// https://leetcode.cn/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				temp := dp[i-1][j]
				if temp > dp[i][j-1] {
					temp = dp[i][j-1]
				}

				dp[i][j] = temp + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	})
}
