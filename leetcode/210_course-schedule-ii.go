package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		graph[from] = append(graph[from], to)
	}

	visited := make([]bool, numCourses)
	paths := make([]bool, numCourses)
	result := make([]int, 0)
	cycle := false

	for i := 0; i < numCourses; i++ {
		traverse(graph, i, visited, paths, &result, &cycle)
		if cycle {
			return []int{}
		}
	}

	reverse210(result)
	return result
}

func findOrderBFS(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0)

	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	count := 0
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		result = append(result, cur)

		for _, v := range graph[cur] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if count == numCourses {
		return result
	}
	return []int{}
}

func reverse210(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func traverse(graph [][]int, idx int, visited, paths []bool, result *[]int, cycle *bool) {
	if paths[idx] {
		*cycle = true
	}

	if visited[idx] || *cycle {
		return
	}

	visited[idx] = true
	paths[idx] = true

	for _, v := range graph[idx] {
		traverse(graph, v, visited, paths, result, cycle)
	}

	*result = append(*result, idx)

	paths[idx] = false
}
