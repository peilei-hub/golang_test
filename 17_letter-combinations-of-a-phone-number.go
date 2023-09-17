package main

import "fmt"

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	digitsMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := make([]string, 0)
	bytes := make([]byte, len(digits))
	backtrace17(digits, digitsMap, &result, 0, bytes)

	return result
}

func backtrace17(digits string, digitsMap map[byte]string, result *[]string, idx int, res []byte) {
	if idx == len(digits) { // 如果idx跟digits长度相同，说明已经遍历到最后一位了，将字符串加到结果里
		*result = append(*result, string(res))
		return
	}

	u := digits[idx]
	str := digitsMap[u]

	// 对可用的选择进行遍历
	for _, b := range str {
		res[idx] = byte(b)                                 // 选择一个
		backtrace17(digits, digitsMap, result, idx+1, res) // 在当前的选择基础上进行下一次递归
		// 回退刚刚的选择
	}
}

func main() {
	combinations := letterCombinations("23")
	fmt.Println(combinations)
}
