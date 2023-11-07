package main

// https://leetcode.cn/problems/set-matrix-zeroes/

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowExist := make([]bool, m)    // 记录每行是否有0
	columnExist := make([]bool, n) // 记录每列是否有0

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				rowExist[i], columnExist[j] = true, true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowExist[i] || columnExist[j] {
				matrix[i][j] = 0
			}
		}
	}
}
