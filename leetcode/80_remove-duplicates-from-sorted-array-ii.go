package main

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	idx := 2 // 需要确定新元素的位置
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[idx-2] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
