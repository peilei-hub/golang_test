package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/restore-ip-addresses/description/

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)

	backtrace93(s, &result, make([]string, 4), 0, 0)

	return result
}

func backtrace93(s string, result *[]string, tmp []string, tmpIdx, start int) {
	if start >= len(s) && tmpIdx == 4 { // 满足条件， start到末尾(遍历完成)，且tmp长度为 4
		join := strings.Join(tmp, ".")
		*result = append(*result, join)
		return
	}

	if tmpIdx >= 4 { // tmpIdx长度大于 4
		return
	}

	// 从 start 开始，长度为 1 到 3 都要遍历。
	for i := start + 1; i <= start+3 && i <= len(s); i++ {
		sub := s[start:i] // 左闭右开

		if isValid93(sub) {
			tmp[tmpIdx] = sub

			backtrace93(s, result, tmp, tmpIdx+1, i)
		}
	}
}

func isValid93(sub string) bool {
	if len(sub) > 1 && sub[0] == '0' {
		return false
	}

	parseInt, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return false
	}

	return parseInt <= 255
}

func main() {
	s := "25525511135"
	//s := "0000"
	addresses := restoreIpAddresses(s)
	fmt.Println(addresses)
}
