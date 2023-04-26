package main

import "math"

// https://leetcode.cn/problems/reverse-integer/

func reverse(x int) int {
	res := 0
	for x != 0 {
		cur := x % 10
		x /= 10

		if res > math.MaxInt/10 || (res == math.MaxInt/10 && cur > 7) {
			return 0
		}
		if res < math.MinInt/10 || (res == math.MinInt/10 && cur < -8) {
			return 0
		}
		res = res*10 + cur
	}

	return res
}
