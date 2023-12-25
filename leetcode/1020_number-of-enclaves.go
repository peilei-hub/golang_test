package main

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		dfs1020(grid, i, 0, m, n)
		dfs1020(grid, i, n-1, m, n)
	}
	for j := 0; j < n; j++ {
		dfs1020(grid, 0, j, m, n)
		dfs1020(grid, m-1, j, m, n)
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				result++
			}
		}
	}
	return result
}

func dfs1020(grid [][]int, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0
	dfs1020(grid, i+1, j, m, n)
	dfs1020(grid, i, j+1, m, n)
	dfs1020(grid, i-1, j, m, n)
	dfs1020(grid, i, j-1, m, n)
}
