package main

// rotate groups of four cells
// time complexity: O(M) where M is the number of cells in the matrix.
// as each cell is getting read once and written once.
//
// space complexity: O(1) because we do not use any other additional data structures.
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}
}

// reverse on diagonal and then reverse left to right
// time complexity: O(M)
// we perform two steps; transposing the matrix, and then reversing each row.
// transposing the matrix has a cost of O(M) because we're moving the value
// of each cell once. reversing each row also has a cost of O(M), because
// again we're moving the value of each cell once.
//
// space complexity: O(1) because we do not use any other additional data structures.
func rotate2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
