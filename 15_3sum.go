package main

import "sort"

// https://leetcode.cn/problems/3sum/

func threeSum(nums []int) [][]int {
	// 长度不到3位直接退出
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums) // 从小到大排序

	// 额外的判断
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return [][]int{{0, 0, 0}}
	}

	res := make([][]int, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > 0 { // 额外判断
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 如果当前位等于前一位，说明已经处理过了
			continue
		}

		// i为选中的一位
		l := i + 1 // l从i下一位开始
		r := n - 1 // r从尾部开始

		// 两边往中间缩小窗口
		for l < r {
			if add := nums[i] + nums[l] + nums[r]; add == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] { // 如果当前位跟前一位相等，则不需要再处理了，跳过重复位，一直到不相等的位置
					l++
				}
				for l < r && nums[r] == nums[r+1] { // 如果当前位跟后一位相等，则不需要再处理了，跳过重复位，一直到不相等的位置
					r--
				}
			} else if add > 0 {
				r--
			} else if add < 0 {
				l++
			}
		}
	}

	return res
}
