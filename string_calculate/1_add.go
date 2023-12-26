package main

import "strconv"

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
