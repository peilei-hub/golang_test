package main

// https://leetcode.cn/problems/all-paths-from-source-to-target/description/

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)

	path := make([]int, 1)
	path[0] = 0

	traverse797(graph, 0, &path, &result)
	return result
}

func traverse797(graph [][]int, idx int, path *[]int, result *[][]int) {
	if idx == len(graph)-1 {
		// 终点
		*result = append(*result, append([]int{}, *path...))
		return
	}

	for _, v := range graph[idx] {
		*path = append(*path, v)
		traverse797(graph, v, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
