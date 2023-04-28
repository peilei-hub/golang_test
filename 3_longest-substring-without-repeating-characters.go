package main

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// method: https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-48c1d/wo-xie-le--f7a92/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	byteSet := make(map[byte]struct{})
	var left, right, max int

	for right < len(s) {
		b := s[right]
		right++

		if _, ok := byteSet[b]; !ok {
			byteSet[b] = struct{}{} // 没有就加进去
		} else {
			for left < right-1 {
				if s[left] != b {
					delete(byteSet, s[left])
					left++
				} else { // 一直找到相等的位置
					left++ // 定位到相等的位置的下一位
					break
				}
			}
		}

		max = intMax(max, len(byteSet))
	}
	return max
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
