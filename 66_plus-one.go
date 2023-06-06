package main

// https://leetcode.cn/problems/plus-one/

func plusOne(digits []int) []int {
	result := make([]int, 0)

	for i := len(digits) - 1; i >= 0; i-- {
		cur := digits[i]

		temp := cur + 1
		if temp < 10 {
			digits[i] = temp
			return digits
		}

		digits[i] = temp % 10
	}

	result = append([]int{1}, digits...)

	return result
}
