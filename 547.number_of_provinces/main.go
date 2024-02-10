package main

// dfs
// time complexity: O(n^2)
// space complexity: O(n)
func findCircleNum(isConnected [][]int) int {
	visited := make(map[int]bool, len(isConnected))
	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true
		for i := 0; i < len(isConnected); i++ {
			if isConnected[node][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	var provinces int
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			provinces++
			dfs(i)
		}
	}

	return provinces
}

// bfs
// time complexity: O(n)
// space complexity: O(n)
func findCircleNum2(isConnected [][]int) int {
	visited := make(map[int]bool, len(isConnected))
	queue := []int{}
	provinces := 0
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			provinces++
			queue = append(queue, i)
		}

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			visited[curr] = true
			for j := 0; j < len(isConnected); j++ {
				if isConnected[curr][j] == 1 && !visited[j] {
					queue = append(queue, j)
				}
			}
		}
	}

	return provinces
}
