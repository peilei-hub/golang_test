package main

// https://leetcode.cn/problems/sqrtx/

func mySqrt(x int) int {
	l := 1
	r := x

	res := 0
	for l <= r {
		mid := l + (r-l)/2
		if mid <= x/mid { // mid * mid <= x
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
