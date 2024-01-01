package main

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/

// 暴力
func strStr1(haystack string, needle string) int {
	len1 := len(haystack)
	len2 := len(needle)

	if len1 == 0 || len2 == 0 || len2 > len1 {
		return -1
	}

	for i := 0; i <= len1-len2; i++ {
		start := i // start

		match := true
		for _, b := range needle {
			if haystack[start] == byte(b) {
				start++
			} else {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/solution/shua-chuan-lc-shuang-bai-po-su-jie-fa-km-tb86/

func main() {
	str1 := "sadbutsad"
	str2 := "sad"

	strStr1(str1, str2)
}
