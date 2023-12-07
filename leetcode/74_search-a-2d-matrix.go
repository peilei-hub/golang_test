package main

// https://leetcode.cn/problems/search-a-2d-matrix/

// O(m + n)

func searchMatrix(matrix [][]int, target int) bool {
	rowMax := len(matrix) - 1
	colMax := len(matrix[0]) - 1

	row, col := 0, colMax

	// 从右上角开始查找
	for row <= rowMax && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target { // 当前值大于target, 向左找
			col--
		} else { // 当前值小于target, 向下找
			row++
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	colMax := len(matrix[0]) - 1

	row := searchRow74(matrix, colMax, target)

	return binarySearch74(matrix[row], target)
}

func binarySearch74(arrays []int, target int) bool {
	left := 0
	right := len(arrays) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arrays[mid] == target {
			return true
		} else if arrays[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func searchRow74(matrix [][]int, colMax int, target int) int { // 查找到可能存在的对应的行号
	left := 0
	right := len(matrix) - 1

	for left <= right {
		mid := left + (right-left)/2
		if num := matrix[mid][colMax]; num == target {
			return mid
		} else if num < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == len(matrix) {
		return left - 1
	}

	return left
}

func main() {
	searchMatrix2([][]int{{1}}, 2)
}
