package main

// BFS
// time complexity: O(n)
// the outer loop ran as long as there were still cells in the queue, dequeuing one each time.
// therefore, it ran at most N times, giving a time complexity of O(N).
//
// the inner loop iterated over the unvisited neighbors of the cell that was dequeued by the outer loop.
// there were at most 8 neighbors. identifying the unvisited neighbors is an O(1) operation because we treat the 8 as a constant.
//
// therefore, we have a time complexity of O(N).
//
// space complexity: O(n)
// the only additional space we used was the queue. we determined above that at most, we enqueued N cells.
// therefore, an upper bound on the worst-case space complexity is O(N).
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid) - 1
	if grid[0][0] != 0 || grid[n][n] != 0 {
		return -1
	}

	directions := [8][2]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	queue := [][]int{{0, 0}}
	grid[0][0] = 1
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		distance := grid[curr[0]][curr[1]]
		if curr[0] == n && curr[1] == n {
			return distance
		}

		for _, dir := range directions {
			x, y := curr[0]+dir[0], curr[1]+dir[1]
			if x < 0 || y < 0 || x > n || y > n || grid[x][y] != 0 {
				continue
			}

			grid[x][y] = distance + 1
			queue = append(queue, []int{x, y})
		}

	}

	return -1
}

// BFS (without overwriting the input)
// time complexity: O(n)
// space complexity: O(n)
func shortestPathBinaryMatrix2(grid [][]int) int {
	n := len(grid) - 1
	if grid[0][0] != 0 || grid[n][n] != 0 {
		return -1
	}

	directions := [8][2]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	queue := [][]int{{0, 0, 1}}
	visited := map[[2]int]bool{
		{0, 0}: true,
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		distance := curr[2]
		if curr[0] == n && curr[1] == n {
			return distance
		}

		for _, dir := range directions {
			x, y := curr[0]+dir[0], curr[1]+dir[1]
			if x < 0 || y < 0 || x > n || y > n || grid[x][y] != 0 || visited[[2]int{x, y}] {
				continue
			}

			visited[[2]int{x, y}] = true
			queue = append(queue, []int{x, y, distance + 1})
		}
	}

	return -1
}
