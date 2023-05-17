package main

import "sort"

// https://leetcode.cn/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	m := make(map[string][]string)
	result := make([][]string, 0)

	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] > bytes[j]
		})

		m[string(bytes)] = append(m[string(bytes)], str)
	}

	for _, strings := range m {
		result = append(result, strings)
	}
	return result
}
