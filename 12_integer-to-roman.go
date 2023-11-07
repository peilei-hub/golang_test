package main

// https://leetcode.cn/problems/integer-to-roman/

func intToRoman(num int) string {
	symbolList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	iList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	for i, v := range iList { // 从最大的开始，一直减到当前数字不能再减，然后用下一位小的数字
		for num >= v {
			num -= v
			result += symbolList[i]
		}
	}

	return result
}
