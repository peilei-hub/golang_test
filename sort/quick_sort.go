package main

import "fmt"

func main() {
	nums := []int{4, 6, 0, 3, 2, 5, 1, 6, 34, 32, 543, 23, 34}
	quickSort(nums, 0, len(nums)-1)

	fmt.Println(nums)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	key := nums[left] // left位置值为空缺

	i, j := left, right

	for i < j {
		// 从右往左找到比 key 小的数
		for i < j && nums[j] >= key {
			j--
		}
		nums[i] = nums[j] // 填补i位置空缺，此时j位置空缺。

		// 从左往右找到比 key 大的数
		for i < j && nums[i] <= key {
			i++
		}
		nums[j] = nums[i] // 填补j位置空缺，此时i位置空缺
	}
	// 此时nums[i]左边都 <= key, nums[i]右边都 >= key
	// 填补i位置空缺
	nums[i] = key

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
