package main

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/3sum-closest/

func threeSumClosest(nums []int, target int) int {
	// 三数之和类似，先排序
	sort.Ints(nums)

	min := math.MaxInt
	result := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // 已经处理过的跳过
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			cur := 0
			if sum == target { // 相等直接返回
				return target
			} else if sum > target { // 大于target,需要将右下标左移，同时记下当前和跟target差值
				cur = sum - target
				r--
			} else if sum < target { // 小于target,需要将左下标左移，同时记下当前和跟target差值
				cur = target - sum
				l++
			}
			if cur < min { // 差值小于记录的最小差值，更新结果
				min = cur
				result = sum
			}
		}

	}
	return result
}
