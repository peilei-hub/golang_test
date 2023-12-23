package main

// https://labuladong.github.io/algo/di-er-zhan-a01c6/bei-bao-le-34bd4/jing-dian--28f3c/

// W weight N个物品

func knapsack(W, N int, wt, val []int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for w := 0; w < W; w++ {
			if w-wt[i-1] < 0 { // 空余容量不够放当前的物品
				// 这种情况下选择不放入背包
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w-wt[i-1]]+val[i-1], // 放入当前物品
					dp[i-1][w]) // 不放入当前物品
			}
		}
	}
	return dp[N][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
