package main

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	bytes := []byte(s)
	byteSet := make(map[byte]struct{})
	var left, right, max int

	for right < len(bytes) {
		b := bytes[right]
		right++

		if _, ok := byteSet[b]; ok {
			for left < right-1 {
				if bytes[left] == b {
					left++
					break
				} else {
					delete(byteSet, bytes[left])
					left++
				}
			}
		} else {
			byteSet[b] = struct{}{}
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
