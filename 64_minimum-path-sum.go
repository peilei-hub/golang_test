package main

// https://leetcode.cn/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range grid {
		dp[i] = make([]int, len(grid[i]))
	}

	for i, row := range grid {
		for j, v := range row {
			if i == 0 && j == 0 {
				dp[i][j] = v
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + v
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + v
			} else {
				tmp := dp[i][j-1]
				if tmp > dp[i-1][j] {
					tmp = dp[i-1][j]
				}
				dp[i][j] = tmp + v
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func main() {
	minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	})
}
