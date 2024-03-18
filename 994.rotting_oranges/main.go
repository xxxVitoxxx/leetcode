package main

func orangesRotting(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	queue := [][]int{}
	freshOrange := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshOrange++
			}
		}
	}

	min := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			curr := queue[0]
			queue = queue[1:]
			for _, d := range directions {
				nextRow, nextCol := curr[0]+d[0], curr[1]+d[1]
				if nextRow >= 0 && nextRow < row && nextCol >= 0 && nextCol < col && grid[nextRow][nextCol] == 1 {
					freshOrange--
					grid[nextRow][nextCol] = 2
					queue = append(queue, []int{nextRow, nextCol})
				}
			}
		}

		if len(queue) > 0 {
			min++
		}
	}

	if freshOrange > 0 {
		return -1
	}

	return min
}
