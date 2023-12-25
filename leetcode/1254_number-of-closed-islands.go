package main

func closedIsland(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ { // 先将边缘的走一遍dfs,将其淹没变成水
		dfs1254(grid, i, 0, m, n)
		dfs1254(grid, i, n-1, m, n)
	}

	for j := 0; j < n; j++ {
		dfs1254(grid, 0, j, m, n)
		dfs1254(grid, m-1, j, m, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs1254(grid, i, j, m, n)
			}
		}
	}
	return res
}

func dfs1254(grid [][]int, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n { // 超过边界
		return
	}

	if grid[i][j] == 1 { // 变成水
		return
	}

	grid[i][j] = 1
	dfs1254(grid, i+1, j, m, n)
	dfs1254(grid, i, j+1, m, n)
	dfs1254(grid, i-1, j, m, n)
	dfs1254(grid, i, j-1, m, n)
}
