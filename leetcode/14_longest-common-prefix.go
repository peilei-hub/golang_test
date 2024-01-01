package main

import "math"

// https://leetcode.cn/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	bytesList := make([][]byte, 0)
	minLen := math.MaxInt
	for _, str := range strs {
		bytes := []byte(str)
		if len(bytes) < minLen {
			minLen = len(bytes)
		}
		bytesList = append(bytesList, bytes)
	}

	result := ""
	for col := 0; col < minLen; col++ {
		var tmp = bytesList[0][col]
		match := true
		for row := 1; row < len(strs); row++ {
			if bytesList[row][col] != tmp {
				match = false
				break
			}
		}
		if !match {
			break
		}

		result += string(tmp)
	}

	return result
}
