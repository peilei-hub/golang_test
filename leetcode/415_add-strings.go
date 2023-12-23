package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	result := ""

	l1 := len(num1) - 1
	l2 := len(num2) - 1

	carry := 0

	for l1 >= 0 || l2 >= 0 {

		var l1Val, l2Val int
		if l1 >= 0 {
			l1Val = int(num1[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			l2Val = int(num2[l2] - '0')
			l2--
		}

		sum := carry + l1Val + l2Val

		tmp := sum % 10
		carry = sum / 10
		result = strconv.FormatInt(int64(tmp), 10) + result
	}
	if carry == 1 {
		result = "1" + result
	}

	return result
}

func subStrings(nums1, nums2 string) string {
	result := ""
	flag := false // 结果是否是负数

	a, b := nums1, nums2 // a为大数，b为小数
	if nums1[0] == '-' {
		a = nums1[1:]
		if len(a) > len(b) || a > b {
			flag = true
		} else {
			a, b = b, a
		}
	} else {
		b = nums2[1:]
		if len(b) > len(a) || b > a {
			flag = true
			a, b = b, a
		}
	}

	borrow := 0
	l1 := len(a) - 1
	l2 := len(b) - 1
	for l1 >= 0 || l2 >= 0 {
		var val1, val2 int
		if l1 >= 0 {
			val1 = int(a[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			val2 = int(b[l2] - '0')
			l2--
		}

		temp := val1 - val2 - borrow
		if temp < 0 {
			borrow = 1
			temp += 10
		} else {
			borrow = 0
		}

		result = strconv.Itoa(temp) + result
	}

	idx := 0
	for i, v := range result {
		if v != '0' {
			idx = i
			break
		}
	}
	result = result[idx:]

	if flag {
		result = "-" + result
	}

	return result
}

func main() {
	strings := subStrings("-110", "30")
	fmt.Println(strings)
}
