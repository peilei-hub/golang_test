package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/minimum-window-substring/

// method: https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-48c1d/wo-xie-le--f7a92/

func minWindow(s string, t string) string {
	target := make(map[byte]int) // 目标包含的每个字符出现的次数
	for _, i := range t {
		target[byte(i)] += 1
	}

	window := make(map[byte]int)
	left, right, validNum := 0, 0, 0

	start := 0
	minLen := math.MaxInt

	for right < len(s) {
		b := s[right]
		right++

		// b 不在目标里，直接跳过
		if _, ok := target[b]; !ok {
			continue
		}

		window[b] += 1
		if target[b] == window[b] { // 相等则有效数加1
			validNum++
		}

		// 处理符合条件的情况，此时要缩小区间
		for validNum == len(target) {
			length := right - left
			if length < minLen {
				start = left
				minLen = length
			}

			leftVal := s[left]
			left++

			if v, ok := window[leftVal]; ok {
				window[leftVal] = v - 1

				if window[leftVal] < target[leftVal] {
					validNum--
				}
			}
		}
	}

	if minLen == math.MaxInt {
		return ""
	}

	return s[start : start+minLen]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	window := minWindow(s, t)
	fmt.Println(window)
}
