package main

import "fmt"

func moveZeroes(nums []int) {
	l, r, n := 0, 0, len(nums)

	for r < n { // l是可以放置非0数据的下标
		if nums[r] != 0 { // 找到非0，就放到l，然后l++
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)
	fmt.Println(nums)
}
