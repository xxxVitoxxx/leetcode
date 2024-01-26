package main

// bfs iteration
// time complexity: O(v+e) where V is the number of the vertices and
// e is the number of the edges.
//
// space complexity: O(v+e)
// use slice record the visited vertices, which takes O(v) space.
// use hash map to store all edges, which O(e) for e edges.
func validPath(edges [][]int, n, source, destination int) bool {
	adjList := make(map[int][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	visited[source] = true
	queue := []int{source}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == destination {
			return true
		}

		for i := 0; i < len(adjList[curr]); i++ {
			if !visited[adjList[curr][i]] {
				visited[adjList[curr][i]] = true
				queue = append(queue, adjList[curr][i])
			}
		}
	}
	return false
}

// dfs iteration
// time complexity: O(v+e)
// space complexity: O(v+e)
func validPath2(edges [][]int, n, source, destination int) bool {
	adjList := make(map[int][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	visited[source] = true
	stack := []int{source}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr == destination {
			return true
		}

		for i := 0; i < len(adjList[curr]); i++ {
			if !visited[adjList[curr][i]] {
				visited[adjList[curr][i]] = true
				stack = append(stack, adjList[curr][i])
			}
		}
	}
	return false
}

// dfs recursion
// time complexity: O(v+e)
// space complexity: O(v+e)
func validPath3(edges [][]int, n, source, destination int) bool {
	adjList := make(map[int][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	visited[source] = true
	var dfs func(int) bool
	dfs = func(curr int) bool {
		if curr == destination {
			return true
		}

		for i := 0; i < len(adjList[curr]); i++ {
			if !visited[adjList[curr][i]] {
				visited[adjList[curr][i]] = true
				if dfs(adjList[curr][i]) {
					return true
				}
			}
		}
		return false
	}
	return dfs(source)
}
