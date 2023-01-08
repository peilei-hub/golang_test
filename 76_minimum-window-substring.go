package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/minimum-window-substring/

// method: https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-48c1d/wo-xie-le--f7a92/

func minWindow(s string, t string) string {
	targetMap := make(map[byte]int)
	windowMap := make(map[byte]int)

	tBytes := []byte(t)
	for _, b := range tBytes {
		if _, ok := targetMap[b]; ok {
			targetMap[b]++
		} else {
			targetMap[b] = 1
		}
	}

	bytes := []byte(s)

	var left, right, start, validNum int
	minLength := math.MaxInt
	for right < len(bytes) {
		b := bytes[right]
		right++

		if _, ok := targetMap[b]; !ok {
			continue
		} else {
			if _, ok := windowMap[b]; ok {
				windowMap[b]++
			} else {
				windowMap[b] = 1
			}
			if windowMap[b] == targetMap[b] {
				validNum++
			}
		}

		for validNum == len(targetMap) {
			length := right - left
			if length < minLength {
				start = left
				minLength = length
			}

			leftVal := bytes[left]
			left++

			if v, ok := windowMap[leftVal]; ok {
				windowMap[leftVal] = v - 1
				if windowMap[leftVal] < targetMap[leftVal] {
					validNum--
				}
			}
		}
	}

	if minLength == math.MaxInt {
		return ""
	}

	return s[start : start+minLength]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	window := minWindow(s, t)
	fmt.Println(window)
}
