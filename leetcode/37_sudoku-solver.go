package main

import "fmt"

// https://leetcode.cn/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	var rowExist, columnExist, blockExist [9][10]bool // 用来存放每行、每列、每块的数字n是否已存在
	for i, row := range board {
		for j, b := range row {
			if b != '.' {
				tmp := b - '0'
				rowExist[i][tmp] = true
				columnExist[j][tmp] = true
				blockExist[i/3*3+j/3][tmp] = true
			}
		}
	}

	selections := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var result [][]byte // 回溯后原本内容会被删除，需要保存下来结果
	for i := 0; i < 9; i++ {
		result = append(result, make([]byte, 9))
	}
	backtrack37(board, rowExist, columnExist, blockExist, 0, 0, selections, result)
	for i, bytes := range result {
		for j, b := range bytes {
			board[i][j] = b
		}
	}
}

func backtrack37(board [][]byte, rowExist, columnExist, blockExist [9][10]bool, rowIdx, columnIdx int, selections []int, result [][]byte) {
	if columnIdx == 9 {
		rowIdx++
		columnIdx = 0
		if rowIdx == 9 { // 需要将结果存起来，因为回溯后会把状态删除
			for i, row := range board {
				for j, b := range row {
					result[i][j] = b
				}
			}
			return
		}
	}

	b := board[rowIdx][columnIdx]
	if b == '.' {
		for _, num := range selections {
			if isValid37(rowExist, columnExist, blockExist, rowIdx, columnIdx, num) { // 遍历所有有效的可能的组合
				board[rowIdx][columnIdx] = byte(num + '0')
				rowExist[rowIdx][num] = true
				columnExist[columnIdx][num] = true
				blockExist[rowIdx/3*3+columnIdx/3][num] = true

				backtrack37(board, rowExist, columnExist, blockExist, rowIdx, columnIdx+1, selections, result)

				board[rowIdx][columnIdx] = '.'
				rowExist[rowIdx][num] = false
				columnExist[columnIdx][num] = false
				blockExist[rowIdx/3*3+columnIdx/3][num] = false
			}
		}
	} else { // 是数字，跳过
		backtrack37(board, rowExist, columnExist, blockExist, rowIdx, columnIdx+1, selections, result)
	}
}

func isValid37(rowExist [9][10]bool, columnExist [9][10]bool, blockExist [9][10]bool, rowIdx int, columnIdx int, num int) bool {
	return !rowExist[rowIdx][num] && !columnExist[columnIdx][num] && !blockExist[rowIdx/3*3+columnIdx/3][num]
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)
	fmt.Println(board)
}
