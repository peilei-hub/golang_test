package main

import "fmt"

// https://leetcode.cn/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	rowMin := 0
	columnMin := 0
	rowMax := len(matrix) - 1
	columnMax := len(matrix[0]) - 1

	result := make([]int, 0)

	for rowMin <= rowMax && columnMin <= columnMax {
		for i := columnMin; i <= columnMax; i++ {
			result = append(result, matrix[rowMin][i])
		}

		for i := rowMin + 1; i <= rowMax; i++ {
			result = append(result, matrix[i][columnMax])
		}

		if rowMin < rowMax {
			for i := columnMax - 1; i >= columnMin; i-- {
				result = append(result, matrix[rowMax][i])
			}
		}

		if columnMin < columnMax {
			for i := rowMax - 1; i > rowMin; i-- {
				result = append(result, matrix[i][columnMin])
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
	order := spiralOrder([][]int{{1, 2, 3, 4}})
	fmt.Println(order)
}
