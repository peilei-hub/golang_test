package main

import "fmt"

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/description/

func search81(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		}

		if nums[m] == nums[l] { // m和l相等无法判断增减就缩小区间
			l++
		} else if nums[m] > nums[l] {
			if binarySearch81(nums, l, m-1, target) {
				return true
			}
			l = m + 1
		} else if nums[m] < nums[l] {
			if binarySearch81(nums, m+1, r, target) {
				return true
			}
			r = m - 1
		}
	}
	return false
}

func binarySearch81(nums []int, left, right, target int) bool {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	return false
}

func main() {
	b := search81([]int{1, 3, 5}, 5)
	fmt.Println(b)
}
