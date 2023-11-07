package main

import "fmt"

// https://leetcode.cn/problems/sort-colors/

func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	i := 0
	for i <= right {
		if nums[i] == 2 { // 如果值为2
			nums[i], nums[right] = nums[right], nums[i] // 跟最右边的替换，此时右边的为2
			right--
		} else if nums[i] == 0 { // 如果值为0
			nums[i], nums[left] = nums[left], nums[i] // 跟最左边的替换，此时左边的为0
			left++

			if i < left { // 从左往右，left++后可能i<left
				i = left
			}
		} else {
			i++
		}
	}
}

func main() {
	tmp := []int{2, 0, 2, 1, 1, 0}
	sortColors(tmp)
	fmt.Println(tmp)
}
