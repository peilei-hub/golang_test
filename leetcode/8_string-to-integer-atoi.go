package main

import "math"

// https://leetcode.cn/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	bytes := []byte(s)

	numType := 1 // 表示正数还是负数
	result := 0
	hasNum := false
	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		if b == ' ' {
			continue
		}

		if b == '-' {
			if i == len(bytes)-1 {
				return 0
			}
			if !isDigit(bytes[i+1]) {
				return 0
			}
			numType = -1
			continue
		}

		if b == '+' {
			if i == len(bytes)-1 {
				return 0
			}
			if !isDigit(bytes[i+1]) {
				return 0
			}
		}

		if isDigit(b) {
			hasNum = true
			num := int(b - '0')
			if numType == 1 {
				if result > math.MaxInt/10 || (result == math.MaxInt/10 && num > 7) {
					return math.MaxInt
				}
			} else {
				num = -num
				if result < math.MinInt/10 || (result == math.MinInt/10 && num < -8) {
					return math.MinInt
				}
			}

			result = result*10 + num

			if i != len(bytes)-1 && !isDigit(bytes[i+1]) {
				return result
			}
		} else if !hasNum { // 从开头找到的第一个（非+ -）字符不是数字
			return 0
		}
	}

	return result
}

func isDigit(b byte) bool {
	return b == '0' || b == '1' || b == '2' || b == '3' || b == '4' || b == '5' || b == '6' || b == '7' || b == '8' || b == '9'
}
