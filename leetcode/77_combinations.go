package main

import "fmt"

// https://leetcode.cn/problems/combinations/

func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	traceback77(&result, n, k, 0, 1, make([]int, k))

	return result
}

func traceback77(result *[][]int, n, k, idx, start int, tmp []int) {
	if idx == k {
		t := make([]int, k)
		copy(t, tmp)
		*result = append(*result, t)
		return
	}

	for i := start; i <= n; i++ {
		tmp[idx] = i

		traceback77(result, n, k, idx+1, i+1, tmp)
	}
}

func main() {
	res := combine(4, 2)
	fmt.Println(res)
}
