package main

// https://leetcode.cn/problems/combinations/

func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	traceback77(&result, make([]int, k), n, k, 0, 1)

	return result
}

func traceback77(result *[][]int, tmp []int, n, k, idx, start int) {
	if idx == k {
		res := make([]int, k)
		for i, v := range tmp {
			res[i] = v
		}
		*result = append(*result, res)
		return
	}

	for i := start; i <= n; i++ {
		tmp[idx] = i

		traceback77(result, tmp, n, k, idx+1, i+1)
	}
}
