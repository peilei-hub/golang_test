package main

// https://leetcode.cn/problems/roman-to-integer/

func romanToInt(s string) int {
	symbolMap := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	result := 0

	for i := 0; i < len(s)-1; i++ {
		cur := symbolMap[s[i]]
		curRight := symbolMap[s[i+1]]

		if cur >= curRight { // 跟右边比较
			result += cur
		} else {
			result -= cur // 如果小于右边，说明当前值需要减去
		}
	}

	result += symbolMap[s[len(s)-1]]

	return result
}
