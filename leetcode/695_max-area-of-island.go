package main

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				max = max695(max, dfs695(grid, i, j, m, n))
			}
		}
	}
	return max
}

func dfs695(grid [][]int, i, j, m, n int) int {
	// 超出索引边界
	if i < 0 || j < 0 || i >= m || j >= n {
		return 0
	}
	// 已经是海水了
	if grid[i][j] == 0 {
		return 0
	}
	// 将 (i, j) 变成海水
	grid[i][j] = 0

	return dfs695(grid, i+1, j, m, n) +
		dfs695(grid, i, j+1, m, n) +
		dfs695(grid, i-1, j, m, n) +
		dfs695(grid, i, j-1, m, n) + 1
}

func max695(a, b int) int {
	if a > b {
		return a
	}
	return b
}
