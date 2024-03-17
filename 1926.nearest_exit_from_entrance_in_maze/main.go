package main

// bfs
// time complexity: O(m * n) in the worst case we may have to visit m * n cells before the iteration stops.
// m stands for width and n stands for height.
// space complexity: O(m + n) we use a queue to store the cells to be visited.
// In the worst case, there may be m + n cells stored in queue.
func nearestExit(maze [][]byte, entrance []int) int {
	row, col := len(maze), len(maze[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{{entrance[0], entrance[1]}}
	// mark the entrance as visited since its not a exit.
	maze[entrance[0]][entrance[1]] = '+'
	step := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			if (curr[0] == 0 || curr[0] == row-1 || curr[1] == 0 || curr[1] == col-1) &&
				(curr[0] != entrance[0] || curr[1] != entrance[1]) {
				return step
			}

			// for the current cell, check its four neighbor cells.
			for _, d := range directions {
				nextRow, nextCol := curr[0]+d[0], curr[1]+d[1]
				if nextRow >= 0 && nextRow < row && nextCol >= 0 && nextCol < col && maze[nextRow][nextCol] == '.' {
					queue = append(queue, []int{nextRow, nextCol})
					maze[nextRow][nextCol] = '+'
				}
			}
		}
		step++
	}

	return -1
}
