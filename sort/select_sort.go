package main

import "fmt"

/**
 * 从数组中选择最小元素，将它与数组的第一个元素交换位置。再从数组剩下的元素中选择出最小的元素，将它与数组的第二个元素交换位置。
 * 不断进行这样的操作，直到将整个数组排序
 * 选择排序，外层一层循环，内层每次找出剩余的最小的值的位置，然后交换 num[i] 和 num[j]
 */

func main() {
	nums := []int{4, 6, 0, 3, 2, 5, 1, 6, 34, 32, 543, 23, 34}
	selectSort(nums)

	fmt.Println(nums)
}

func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		nums[minIdx], nums[i] = nums[i], nums[minIdx]
	}
}
