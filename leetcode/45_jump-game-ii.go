package main

import "fmt"

// https://leetcode.cn/problems/jump-game-ii/

func jump(nums []int) int {
	steps := 0

	nextBorder := 0 // 下一个边界
	curBorder := 0  // 当前能跳到的边界

	// i表示从哪个位置起跳，所以 i 需要<len(nums)-1。因为到最后一位就不需要跳了
	for i := 0; i < len(nums)-1; i++ {
		nextBorder = max45(nextBorder, nums[i]+i) // 在当前border内去找下次的border
		if i == curBorder {                       // 到达当前border
			curBorder = nextBorder
			steps++
		}
	}
	return steps
}

func max45(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	i := jump([]int{2, 3, 1, 1, 4})
	fmt.Println(i)
}
