package main

// https://leetcode.cn/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	var rowExist, columnExist, blockExist [9][10]bool // 用来存放每行、每列、每块的数字n是否已存在

	for i, rows := range board {
		for j, b := range rows {
			if b != '.' {
				tmp := b - '0'
				rowExist[i][tmp] = true
				columnExist[j][tmp] = true
				blockExist[i/3*3+j/3][tmp] = true
			}
		}
	}

	selectedNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var result [][]byte
	for i := 0; i < 9; i++ {
		result = append(result, make([]byte, 9))
	}

	// 从row遍历，每次column+1，判断边界值和终止值
	backtrace37(selectedNums, &board, 0, 0, &rowExist, &columnExist, &blockExist, &result)

	for i, bytes := range result {
		for j, b := range bytes {
			board[i][j] = b
		}
	}
}

func backtrace37(selectedNums []int, board *[][]byte, row, column int, rowExist, columnExist, blockExist *[9][10]bool, result *[][]byte) {
	if column == 9 {
		column = 0
		row++
		if row == 9 {
			// 需要将结果存起来，因为回溯后会把状态删除
			for i, bytes := range *board {
				for j, b := range bytes {
					(*result)[i][j] = b
				}
			}
			return
		}
	}

	rows := rowExist[row]
	columns := columnExist[column]
	blocks := blockExist[row/3*3+column/3]

	b := (*board)[row][column]
	if b == '.' { // 不存在数字
		for _, num := range selectedNums { // 需要对所有可能的结果进行遍历选择
			if isValid37(rows, columns, blocks, num) {
				(*board)[row][column] = byte(num + '0')
				rowExist[row][num] = true
				columnExist[column][num] = true
				blockExist[row/3*3+column/3][num] = true

				backtrace37(selectedNums, board, row, column+1, rowExist, columnExist, blockExist, result)

				(*board)[row][column] = '.'
				rowExist[row][num] = false
				columnExist[column][num] = false
				blockExist[row/3*3+column/3][num] = false
			}
		}
	} else { // 当前已有数字，直接下一个格子
		backtrace37(selectedNums, board, row, column+1, rowExist, columnExist, blockExist, result)
	}
}

func isValid37(rowExist, columnExist, blockExist [10]bool, num int) bool {
	return !rowExist[num] && !columnExist[num] && !blockExist[num]
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
}
