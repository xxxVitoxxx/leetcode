package main

// time complexity: O(n*m)
// space complexity: O(n*m)
func zigzagTraversal(grid [][]int) []int {
	result := make([]int, 0, (len(grid)*len(grid[0])+1)>>1)
	colLen := len(grid[0])
	toRightInitIdx, toLeftInitIdx := 0, colLen-1
	if colLen%2 == 1 {
		toLeftInitIdx = colLen - 2
	}

	interval := 2
	for row := 0; row < len(grid); row++ {
		col := toRightInitIdx
		if row%2 == 1 {
			col = toLeftInitIdx
		}

		for ; col >= 0 && col < colLen; col += interval {
			result = append(result, grid[row][col])
		}
		interval *= -1
	}

	return result
}

// time complexity: O(n*m)
// space complexity: O(n*m)
func zigzagTraversal2(grid [][]int) []int {
	result := make([]int, 0, (len(grid)*len(grid[0])+1)>>1)
	skip := false
	step := 1
	for row, col := 0, 0; row < len(grid); row++ {
		for ; col >= 0 && col < len(grid[0]); col += step {
			if !skip {
				result = append(result, grid[row][col])
			}
			skip = !skip
		}
		col -= step
		step *= -1
	}
	return result
}

// time complexity: O(n*m)
// space complexity: O(n*m)
func zigzagTraversal3(grid [][]int) []int {
	result := make([]int, 0, (len(grid)*len(grid[0])+1)>>1)
	for row := 0; row < len(grid); row++ {
		if row%2 == 0 {
			for col := 0; col < len(grid[0]); col++ {
				if col%2 == 0 {
					result = append(result, grid[row][col])
				}
			}
		} else {
			for col := len(grid[0]) - 1; col >= 0; col-- {
				if col%2 == 1 {
					result = append(result, grid[row][col])
				}
			}
		}
	}
	return result
}
