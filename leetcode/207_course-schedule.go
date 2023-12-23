package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses) // 构造图
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		graph[from] = append(graph[from], to)
	}

	visited := make([]bool, numCourses) // 是否遍历过
	onPath := make([]bool, numCourses)  // 是否在路径上
	cycle := false
	for i := 0; i < numCourses; i++ {
		traverse207(graph, i, visited, onPath, &cycle)
	}

	return !cycle
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

func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses) // 构造图
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	inDegree := make([]int, numCourses) // 入度数组
	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	queue := make([]int, 0) // 存放入度为0的节点
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++                        // 遍历到的节点数量++
		for _, v := range graph[cur] { // 遍历当前节点的连接节点，将对应节点的入度--
			inDegree[v]--
			if inDegree[v] == 0 { // 如果为0，放到queue里
				queue = append(queue, v)
			}
		}
	}

	return count == numCourses
}
