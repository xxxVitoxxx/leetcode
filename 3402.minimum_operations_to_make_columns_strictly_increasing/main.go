package main

// time complexity: O(n*m)
// space complexity: O(1)
func minimumOperations(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	result := 0
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] <= grid[i-1][j] {
				temp := grid[i][j]
				grid[i][j] = grid[i-1][j] + 1
				result += grid[i][j] - temp
			}
		}
	}

	return result
}
