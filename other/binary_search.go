package main

import "fmt"

func binarySearch(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			left = mid + 1
		} else if nums[mid] < target {
			right = mid - 1
		}
	}

	return false
}

func leftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1 // 右区间缩小
		} else if nums[mid] > target {
			left = mid + 1
		} else if nums[mid] < target {
			right = mid - 1
		}
	}

	return -1
}

func rightBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			left = mid + 1 // 左区间缩小
		} else if nums[mid] > target {
			left = mid + 1
		} else if nums[mid] < target {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	bound := leftBound([]int{1, 2, 3}, 2)
	fmt.Println(bound)
}
