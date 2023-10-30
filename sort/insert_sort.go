package main

import "fmt"

func main() {
	nums := []int{4, 6, 0, 3, 2, 5, 1, 6, 34, 32, 543, 23, 34}
	insertSort(nums)

	fmt.Println(nums)
}

/**
 * 每次只交换相邻的元素
 * 每次都将当前元素插入到左侧已经排序的数组中，使得插入之后左侧数组依然有序
 * 插入先认为第一个数有序，从第二个数开始循环。内层循环为逆循环，将依次往前查找，并交换位置
 * 类似每次从桌面拿一张扑克牌，然后插入到有序的数组中
 */
func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ { // 0 为有序，所以从下标 1 开始
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
