package main

// https://leetcode.cn/problems/rotate-image/

func rotate(matrix [][]int) {
	// 找规律，先转置，再左右翻转

	for i, row := range matrix {
		for j := range row {
			if j > i {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}

	for _, row := range matrix {
		i := 0
		j := len(row) - 1
		for i < j {
			row[i], row[j] = row[j], row[i]
			i++
			j--
		}
	}
}
