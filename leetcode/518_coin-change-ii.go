package main

// https://leetcode.cn/problems/coin-change-ii

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	// for coin in coins循环中就把coin的可使用次数规定好了。
	// 不允许在后面的硬币层次使用之前的硬币。
	// 这就像排列中2,2,1; 2,1,2是两种情况，
	// 但是组合问题规定好了一种书写顺序，比如大的写在前面那就只有2,2,1这一种情况了。
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
