package main

// DFS
// time complexity: O(M*N)
// space complexity: O(M*N)
func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific, atlantic := make([][]bool, m), make([][]bool, m)
	for i := range m {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var dfs func(int, int, [][]bool)
	dfs = func(row, col int, visited [][]bool) {
		visited[row][col] = true

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]

			if newRow < 0 || newCol < 0 || newRow >= m || newCol >= n {
				continue
			}

			if visited[newRow][newCol] || heights[newRow][newCol] < heights[row][col] {
				continue
			}

			dfs(newRow, newCol, visited)
		}
	}

	for i := range m {
		dfs(i, 0, pacific)
		dfs(i, n-1, atlantic)
	}

	for i := range n {
		dfs(0, i, pacific)
		dfs(m-1, i, atlantic)
	}

	var result [][]int
	for i := range m {
		for j := range n {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// BFS
// time complexity: O(M*N)
// space complexity: O(M*N)
func pacificAtlantic2(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacificQueue, atlanticQueue := [][2]int{}, [][2]int{}
	for i := range m {
		pacificQueue = append(pacificQueue, [2]int{i, 0})
		atlanticQueue = append(atlanticQueue, [2]int{i, n - 1})
	}

	for i := range n {
		pacificQueue = append(pacificQueue, [2]int{0, i})
		atlanticQueue = append(atlanticQueue, [2]int{m - 1, i})
	}

	pacific := bfs(heights, pacificQueue)
	atlantic := bfs(heights, atlanticQueue)

	var result [][]int
	for i := range m {
		for j := range n {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func bfs(heights [][]int, queue [][2]int) [][]bool {
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	visited := make([][]bool, len(heights))
	for i := range len(heights) {
		visited[i] = make([]bool, len(heights[i]))
	}

	for _, q := range queue {
		visited[q[0]][q[1]] = true
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow, newCol := curr[0]+dir[0], curr[1]+dir[1]
			if newRow < 0 || newCol < 0 || newRow >= len(heights) || newCol >= len(heights[0]) {
				continue
			}

			if visited[newRow][newCol] || heights[curr[0]][curr[1]] > heights[newRow][newCol] {
				continue
			}

			visited[newRow][newCol] = true
			queue = append(queue, [2]int{newRow, newCol})
		}
	}

	return visited
}
