package main

// bfs
// time complexity: O(n * m) where n * m is the size of the grid
// space complexity: O(n * m) in the worst case, the grid is filled with rotten oranges.
// as a result, the queue would the initialized with all the cells in the grid.

// normally for BFS, the main space complexity lies in process rather than the initialization.
// for instance, for a BFS traversal in a binary tree, at any given moment,
// the queue would hold no more than 2 levels of tree nodes.
// therefore, the space complexity of BFS traversal in a binary tree would depend on the width of the input tree.
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

// in-place bfs
// time complexity: O(n * m)
// space complexity: O(1)
func orangesRotting2(grid [][]int) int {
	moreToRot := true
	time := 2
	for moreToRot {
		moreToRot = rottenProcess(grid, time)
		if moreToRot {
			time++
		}
	}

	if allRotten(grid) {
		return time - 2
	}
	return -1
}

func rottenProcess(grid [][]int, time int) bool {
	var moreToRot bool
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == time {
				moreToRot = neighborToRotten(grid, time+1, i, j) || moreToRot
			}
		}
	}
	return moreToRot
}

func neighborToRotten(grid [][]int, time, row, col int) bool {
	var moreToRot bool
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, d := range directions {
		moreToRot = rottenIfFresh(grid, time, row+d[0], col+d[1]) || moreToRot
	}
	return moreToRot
}

func rottenIfFresh(grid [][]int, time, row, col int) bool {
	if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == 1 {
		grid[row][col] = time
		return true
	}
	return false
}

func allRotten(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return false
			}
		}
	}
	return true
}
