package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // 从0到n

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		curMin := math.MaxInt32
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			if dp[i-coin] == math.MaxInt32 {
				continue
			}

			curMin = min322(curMin, 1+dp[i-coin])
		}
		dp[i] = curMin
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min322(a, b int) int {
	if a > b {
		return b
	}
	return a
}
