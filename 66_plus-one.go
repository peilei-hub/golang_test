package main

// https://leetcode.cn/problems/plus-one/

func plusOne(digits []int) []int {
	result := make([]int, 0)

	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + 1
		if tmp < 10 {
			digits[i] = tmp
			return digits
		}
		digits[i] = tmp % 10
	}
	result = append([]int{1}, digits...)

	return result
}
