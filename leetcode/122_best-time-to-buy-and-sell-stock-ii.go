package main

func maxProfit2(prices []int) int {
	days := len(prices)
	dp := make([][]int, days) // 第i天是否持有
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < days; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		dp[i][0] = max122(dp[i-1][0], dp[i-1][1]+prices[i]) // i-1天未持有，或者i-1天持有i天卖出
		dp[i][1] = max122(dp[i-1][1], dp[i-1][0]-prices[i]) // i-1天持有，或者i天买入
	}

	return dp[days-1][0]
}

func max122(a, b int) int {
	if a > b {
		return a
	}
	return b
}
