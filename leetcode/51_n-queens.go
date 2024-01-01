package main

import "fmt"

// https://leetcode.cn/problems/n-queens/

func solveNQueens(n int) [][]string {
	tables := make([][]byte, 0, n) //  初始化
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		tables = append(tables, row)
	}

	result := make([][]string, 0)

	// 每次填新的一行
	backtrack51(&result, n, 0, tables)

	return result
}

func backtrack51(result *[][]string, n int, rowIdx int, tables [][]byte) {
	if rowIdx == n {
		tmp := make([]string, 0)
		for _, table := range tables {
			tmp = append(tmp, string(table))
		}

		*result = append(*result, tmp)
		return
	}

	for columnIdx := 0; columnIdx < n; columnIdx++ {
		if isValid51(tables, rowIdx, columnIdx) {
			tables[rowIdx][columnIdx] = 'Q'

			backtrack51(result, n, rowIdx+1, tables)

			tables[rowIdx][columnIdx] = '.'
		}
	}
}

func isValid51(tables [][]byte, rowIdx, columnIdx int) bool {
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

func main() {
	queens := solveNQueens(4)
	fmt.Println(queens)
}
