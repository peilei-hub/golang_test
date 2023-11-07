package main

import "fmt"

func main() {
	nums := []int{4, 6, 0, 3, 2, 5, 1, 6, 34, 32, 543, 23, 34}
	bubbleSort(nums)

	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- { // 每次将最大的数放到最后。下一次遍历只需遍历从0到上次最大的数之前的位置
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
