package main

import "fmt"

func moveZeroes(nums []int) {
	idx, i := 0, 0

	for i <= len(nums)-1 { // idx是可以放置非0数据的下标
		if nums[i] != 0 { // 找到非0，就放到idx，然后idx++
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
		i++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)
	fmt.Println(nums)
}
