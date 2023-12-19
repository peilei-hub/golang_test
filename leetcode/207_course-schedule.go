package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		graph[from] = append(graph[from], to)
	}

	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false
	for i := 0; i < numCourses; i++ {
		traverse207(graph, i, visited, onPath, &hasCycle)
	}

	return !hasCycle
}

func traverse207(graph [][]int, i int, visited, onPath []bool, hasCycle *bool) {
	if onPath[i] {
		*hasCycle = true
	}
	if visited[i] || *hasCycle {
		return
	}

	visited[i] = true
	onPath[i] = true
	for _, v := range graph[i] {
		traverse207(graph, v, visited, onPath, hasCycle)
	}
	onPath[i] = false
}
