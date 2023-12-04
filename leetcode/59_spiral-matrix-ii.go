package main

import "fmt"

// https://leetcode.cn/problems/spiral-matrix-ii/

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	rowMin, columnMin, rowMax, columnMax := 0, 0, n-1, n-1
	temp := 1

	for rowMin <= rowMax && columnMin <= columnMax {
		for i := columnMin; i <= columnMax; i++ {
			result[rowMin][i] = temp
			temp++
		}

		for i := rowMin + 1; i <= rowMax; i++ {
			result[i][columnMax] = temp
			temp++
		}

		if rowMin < rowMax {
			for i := columnMax - 1; i >= columnMin; i-- {
				result[rowMax][i] = temp
				temp++
			}
		}

		if columnMin < columnMax {
			for i := rowMax - 1; i > rowMin; i-- {
				result[i][columnMin] = temp
				temp++
			}
		}

		rowMin++
		columnMin++
		rowMax--
		columnMax--
	}

	return result
}

func main() {
	matrix := generateMatrix(10)
	fmt.Println(matrix)
}
