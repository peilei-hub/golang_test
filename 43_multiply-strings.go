package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/multiply-strings/

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
			result = addStr43(tmpStr, result)

			innerZeros += "0"
		}

		outerZeros += "0"
	}

	return result
}

func addStr43(nums1 string, nums2 string) string {
	result := ""

	carry := 0

	len1 := len(nums1) - 1
	len2 := len(nums2) - 1

	for len1 >= 0 || len2 >= 0 {
		var v1, v2 int
		if len1 >= 0 {
			v1 = int(nums1[len1] - '0')
			len1--
		}
		if len2 >= 0 {
			v2 = int(nums2[len2] - '0')
			len2--
		}

		tmp := v1 + v2 + carry

		v := tmp % 10
		carry = tmp / 10

		result = strconv.FormatInt(int64(v), 10) + result
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}

func main() {
	s := multiply("123", "456")
	fmt.Println(s)
}
