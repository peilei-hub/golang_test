package main

// https://leetcode.cn/problems/roman-to-integer/

func romanToIntV1(s string) int {
	symbolMap := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			if val, ok := symbolMap[s[i:i+2]]; ok {
				result += val
				i++
			} else {
				result += symbolMap[s[i:i+1]]
			}
		} else {
			result += symbolMap[s[i:i+1]]
		}
	}

	return result
}

func romanToIntV2(s string) int {
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
