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

	key := nums[left]

	i, j := left, right

	for i < j {
		// 从右往左找到比 key 小的数
		for i < j && nums[j] >= key {
			j--
		}

		// 从左往右找到比 key 大的数
		for i < j && nums[i] <= key {
			i++
		}

		// 交换
		if i < j {
			swap(nums, i, j)
		}
	}

	// i = j
	swap(nums, i, left)
	//nums[left] = nums[i]
	//nums[i] = key

	quickSort(nums, left, i-1)
	quickSort(nums, j+1, right)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
