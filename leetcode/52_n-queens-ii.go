package main

// https://leetcode.cn/problems/n-queens-ii/

func totalNQueens(n int) int {

	tables := make([][]byte, 0, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		tables = append(tables, row)
	}

	result := 0

	traceback52(tables, &result, 0, n)

	return result
}

func traceback52(tables [][]byte, result *int, rowIdx int, n int) {
	if rowIdx == n {
		*result++
		return
	}

	for columnIdx := 0; columnIdx < n; columnIdx++ {
		if isValid52(tables, rowIdx, columnIdx) {
			tables[rowIdx][columnIdx] = 'Q'

			traceback52(tables, result, rowIdx+1, n)

			tables[rowIdx][columnIdx] = '.'
		}
	}
}

func isValid52(tables [][]byte, rowIdx, columnIdx int) bool {
	// 上方
	for i := 0; i < rowIdx; i++ {
		if tables[i][columnIdx] == 'Q' {
			return false
		}
	}

	// 左上
	for i, j := rowIdx-1, columnIdx-1; i >= 0 && j >= 0; {
		if tables[i][j] == 'Q' {
			return false
		}

		i--
		j--
	}

	// 右上
	for i, j := rowIdx-1, columnIdx+1; i >= 0 && j < len(tables); {
		if tables[i][j] == 'Q' {
			return false
		}

		i--
		j++
	}

	return true
}
