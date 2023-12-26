package main

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := "0"

	// 一个乘数的位数后面应该带的0
	outerZeros := ""
	for i := len(num1) - 1; i >= 0; i-- {

		// 另一个乘数的位数后面应该带的0
		innerZeros := ""
		for j := len(num2) - 1; j >= 0; j-- {
			tmp := int(num1[i]-'0') * int(num2[j]-'0')

			tmpStr := strconv.FormatInt(int64(tmp), 10)
			tmpStr += outerZeros + innerZeros // 补0
			result = addStrings(tmpStr, result)

			innerZeros += "0"
		}

		outerZeros += "0"
	}

	return result
}
