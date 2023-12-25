package main

// https://leetcode.cn/problems/number-of-islands

func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m) // 使用visited
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] { // 使用visited
				res++
				dfs200(grid, i, j, m, n, visited)
			}
		}
	}
	return res
}

func dfs200(grid [][]byte, i, j, m, n int, visited [][]bool) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if visited[i][j] {
		return
	}
	if grid[i][j] == '0' {
		return
	}

	visited[i][j] = true
	dfs200(grid, i+1, j, m, n, visited)
	dfs200(grid, i, j+1, m, n, visited)
	dfs200(grid, i-1, j, m, n, visited)
	dfs200(grid, i, j-1, m, n, visited)
}

// 不使用visited
func numIslands2(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++                     // 每发现一个岛屿，岛屿数量加一
				dfs2002(grid, i, j, m, n) // 然后使用DFS将岛屿淹了
			}
		}
	}
	return res
}

func dfs2002(grid [][]byte, i, j, m, n int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}

	// 将 (i,j) 变成海水
	grid[i][j] = '0'
	dfs2002(grid, i+1, j, m, n)
	dfs2002(grid, i, j+1, m, n)
	dfs2002(grid, i-1, j, m, n)
	dfs2002(grid, i, j-1, m, n)
}
