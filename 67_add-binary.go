package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/add-binary/

func addBinary(a string, b string) string {
	result := ""

	i, j := len(a)-1, len(b)-1

	carry := 0
	for i >= 0 || j >= 0 {
		iVal, jVal := 0, 0
		if i >= 0 {
			iVal = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			jVal = int(b[j] - '0')
			j--
		}

		temp := iVal + jVal + carry
		carry = temp / 2
		temp %= 2

		result = strconv.Itoa(temp) + result
	}

	if carry == 1 {
		result = "1" + result
	}
	return result
}

func main() {
	binary := addBinary("1010", "1011")
	fmt.Println(binary)
}
