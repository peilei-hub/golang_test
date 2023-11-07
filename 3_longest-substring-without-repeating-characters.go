package main

import "wangpeilei/leetcode/utils"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// method: https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-48c1d/wo-xie-le--f7a92/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	byteSet := make(map[byte]struct{})
	var left, right, max int

	for right < len(s) {
		b := s[right] // 固定写法
		right++

		if _, ok := byteSet[b]; !ok {
			byteSet[b] = struct{}{} // 没有就加进去
		} else { // set里已有，需要找到之前加入的这个字符的位置
			for left < right-1 {
				leftVal := s[left] // 固定写法
				left++

				if b != leftVal { // 不相等
					delete(byteSet, leftVal)
				} else { // 一直到相等的边界
					break
				}
			}
		}

		max = utils.Max(max, len(byteSet))
	}
	return max
}
