package main

import "sort"

// https://leetcode.cn/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l := j + 1
			r := len(nums) - 1
			for l < r {
				if add := nums[i] + nums[j] + nums[l] + nums[r]; add == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if add > target {
					r--
				} else if add < target {
					l++
				}
			}
		}
	}

	return res
}
