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
	for i := 0; i < minLen; i++ {
		var tmp = bytesList[0][i]
		match := true
		for j := 1; j < len(strs); j++ {
			if bytesList[j][i] != tmp {
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
