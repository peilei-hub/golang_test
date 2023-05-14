package main

// https://leetcode.cn/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	var rowExist, columnExist, blockExist [9][10]bool

	for i, rows := range board {
		for j, b := range rows {
			if b != '.' {
				tmp := b - '0'
				if rowExist[i][tmp] {
					return false
				}
				rowExist[i][tmp] = true

				if columnExist[j][tmp] {
					return false
				}
				columnExist[j][tmp] = true

				if blockExist[i/3*3+j/3][tmp] {
					return false
				}
				blockExist[i/3*3+j/3][tmp] = true
			}
		}
	}

	return true
}
