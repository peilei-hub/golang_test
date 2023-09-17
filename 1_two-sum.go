package main

// https://leetcode.cn/problems/two-sum/

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int) // key 为 num, value 为 num 对应的下标
	for i, num := range nums {
		if v, ok := numMap[target-num]; ok {
			return []int{i, v}
		}
		numMap[num] = i
	}
	return nil
}
