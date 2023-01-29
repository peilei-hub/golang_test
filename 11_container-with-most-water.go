package main

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		max := minInt(height[l], height[r]) * (l - r)
		if max > res {
			res = max
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return res
}

func minInt(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
