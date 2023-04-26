package main

import "sort"

// https://leetcode.cn/problems/3sum/

func threeSum(nums []int) [][]int {
	// 长度不到3位直接退出
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums) // 从小到大排序
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return [][]int{{0, 0, 0}}
	}

	res := make([][]int, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := n - 1

		for l < r {
			if add := nums[i] + nums[l] + nums[r]; add == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
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
