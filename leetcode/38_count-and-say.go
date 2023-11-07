package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/count-and-say/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := countAndSay(n - 1) // 找出n-1的str

	result := ""

	// 然后对n-1的str进行描述
	var count int
	for i := range str {
		if i == 0 {
			count = 1
			continue
		}

		if str[i] == str[i-1] {
			count++
		} else {
			result += strconv.FormatInt(int64(count), 10) + string(str[i-1]) // 出现的次数 + 值
			count = 1
		}
	}

	result += strconv.FormatInt(int64(count), 10) + string(str[len(str)-1]) // 最后一个数出现的次数 + 值

	return result
}

func main() {
	say := countAndSay(4)
	fmt.Println(say)
}
