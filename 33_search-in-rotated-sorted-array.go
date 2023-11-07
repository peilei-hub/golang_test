package main

// https://leetcode.cn/problems/search-in-rotated-sorted-array/

// 思路：始终找出升序的部分进行搜索

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target { // 如果等于直接返回
			return m
		}

		if nums[m] > nums[l] { // 中间大于左边，说明从左到中间是递增，直接二分查找
			i := binarySearch(nums, l, m-1, target)
			if i != -1 {
				return i
			} else {
				l = m + 1 // 左边查过未找到, 将l=m+1，下一轮会找右边的
			}
		} else { // 中间小于左边，说明从中间到右是递增，二分查找右边的
			i := binarySearch(nums, m+1, r, target)
			if i != -1 {
				return i
			} else {
				r = m - 1 // 右边未找到，将r=m-1，下一轮会找左边的
			}
		}
	}

	return -1
}

func binarySearch(nums []int, l, r, target int) int {
	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		}
	}

	return -1
}
