package main

import "fmt"

// https://leetcode.cn/problems/jump-game-ii/

func jump(nums []int) int {
	steps := 0

	nextBorder := 0
	border := 0

	for i := 0; i < len(nums)-1; i++ {
		nextBorder = max45(nextBorder, nums[i]+i) // 在当前border内去找下次的border
		if i == border {                          // 到答当前border
			steps++
			border = nextBorder
		}
	}
	return steps
}

//func jump(nums []int) int {
//	steps := 0
//
//	border := 0 // 维护当前能够到达的最大下标位置，记为边界
//	end := 0
//	for i := 0; i < len(nums)-1; i++ {
//		border = max45(border, nums[i]+i) // 在当次到达边界时，每次都找到最远的当下次的边界
//		if i == end {                   // 到达当次边界
//			end = border
//			steps++ // 在到达边界之内只会跳一次，所以到达边界再更新step
//		}
//	}
//
//	return steps
//}

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
