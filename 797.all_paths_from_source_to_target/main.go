package main

// DFS
// time complexity: O(2^n * n)
// space complexity: O(n)
func allPathsSourceTarget(graph [][]int) [][]int {
	adjList := make(map[int][]int)
	for i, nbs := range graph {
		adjList[i] = nbs
	}

	var paths [][]int
	var path []int
	var dfs func(int)
	dfs = func(n int) {
		path = append(path, n)

		if n == len(graph)-1 {
			p := make([]int, len(path))
			copy(p, path)
			paths = append(paths, p)
			return
		}

		for _, nb := range adjList[n] {
			dfs(nb)
			path = path[:len(path)-1]
		}
	}

	dfs(0)

	return paths
}

// BFS
func allPathsSourceTarget2(graph [][]int) [][]int {
	adjList := make(map[int][]int)
	for i, nbs := range graph {
		adjList[i] = nbs
	}

	paths := [][]int{}
	queue := [][]int{{0}}
	for len(queue) > 0 {
		currPath := queue[0]
		queue = queue[1:]

		vertex := currPath[len(currPath)-1]
		for _, nb := range adjList[vertex] {
			temp := make([]int, len(currPath))
			copy(temp, currPath)

			temp = append(temp, nb)

			if nb == len(graph)-1 {
				paths = append(paths, temp)
			} else {
				queue = append(queue, temp)
			}
		}
	}

	return paths
}
