package main

import "fmt"

func numTrees(n int) int {
	if n == 0 {
		return 0
	}

	return allTrees2(1, n)
}

// 获得 [start,end]的树列表
func allTrees2(start, end int) int {
	if start > end {
		return 0
	}
	if start == end {
		return 1
	}

	result := 0

	for i := start; i <= end; i++ {
		leftTrees := allTrees2(start, i-1)
		rightTrees := allTrees2(i+1, end)

		if leftTrees == 0 {
			if rightTrees != 0 {
				result += rightTrees
			}
		} else {
			if rightTrees == 0 {
				result += leftTrees
			} else {
				result += leftTrees * rightTrees
			}
		}

	}

	return result
}

func main() {
	trees := numTrees(18)
	fmt.Println(trees)
}
