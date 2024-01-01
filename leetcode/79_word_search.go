package main

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	columnLen := len(board[0])
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, columnLen)
	}

	for i, row := range board {
		for j, b := range row {
			if b == word[0] && backtrack79(board, visited, i, j, len(board), columnLen, 0, word) {
				return true
			}
		}
	}

	return false
}

func backtrack79(board [][]byte, visited [][]bool, i, j, rowLen, columnLen, idx int, word string) bool {
	if idx == len(word) {
		return true
	}
	if board[i][j] != word[idx] {
		return false
	}

	// board[i][j] = word[idx] 的情况

	if idx == len(word)-1 {
		return true
	}

	visited[i][j] = true // 置为访问
	res := false
	if j-1 >= 0 && !visited[i][j-1] { // 左边
		if backtrack79(board, visited, i, j-1, rowLen, columnLen, idx+1, word) {
			res = true
		}
	}
	if !res && j+1 <= columnLen-1 && !visited[i][j+1] { // 右边
		if backtrack79(board, visited, i, j+1, rowLen, columnLen, idx+1, word) {
			res = true
		}
	}
	if !res && i-1 >= 0 && !visited[i-1][j] { // 上方
		if backtrack79(board, visited, i-1, j, rowLen, columnLen, idx+1, word) {
			res = true
		}
	}
	if !res && i+1 <= rowLen-1 && !visited[i+1][j] { // 下方
		if backtrack79(board, visited, i+1, j, rowLen, columnLen, idx+1, word) {
			res = true
		}
	}

	visited[i][j] = false // 置为未访问
	return res
}
