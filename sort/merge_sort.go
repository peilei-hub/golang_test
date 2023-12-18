package main

import "fmt"

func main() {
	nums := []int{4, 6, 0, 3, 2, 5, 1, 6, 34, 32, 543, 23, 34}
	mergeSort(nums)

	fmt.Println(nums)
}

func mergeSort(nums []int) {
	mergeSortInner(nums, 0, len(nums)-1, make([]int, len(nums))) // 用一个辅助数组暂存部分归并的数据
}

func mergeSortInner(nums []int, left, right int, temp []int) { // 将 [left,right]有序
	if left < right {
		mid := left + (right-left)/2
		mergeSortInner(nums, left, mid, temp)
		mergeSortInner(nums, mid+1, right, temp)
		merge(nums, left, mid, right, temp)
	}
}

func merge(nums []int, left, mid, right int, temp []int) {
	// [left,mid]
	// [mid+1,right]
	idx := 0

	leftBegin := left
	rightBegin := mid + 1
	for leftBegin <= mid && rightBegin <= right {
		if nums[leftBegin] < nums[rightBegin] {
			temp[idx] = nums[leftBegin]
			idx++
			leftBegin++
		} else {
			temp[idx] = nums[rightBegin]
			idx++
			rightBegin++
		}
	}

	for leftBegin <= mid {
		temp[idx] = nums[leftBegin]
		idx++
		leftBegin++
	}

	for rightBegin <= right {
		temp[idx] = nums[rightBegin]
		idx++
		rightBegin++
	}

	idx = 0
	for left <= right {
		nums[left] = temp[idx]
		left++
		idx++
	}

}
