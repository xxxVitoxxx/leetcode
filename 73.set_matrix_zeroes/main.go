package main

// additional memory approach
// time complexity: O(M * N) where M and N are the number
// of the rows and columns
//
// space complexity: O(M + N)
func setZeroes(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i], cols[j] = true, true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// use the matrix itself as the indicators
// time complexity: O(M * N)
// space complexity: O(1)
//
// the ideas is that we can use the first cell of every row
// and column as a flag. this flag would determine whether
// a row or column has been set to zero.
func setZeroes2(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	zeroRow, zeroCol := false, false
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				if row == 0 {
					zeroRow = true
				} else {
					matrix[row][0] = 0
				}

				if col == 0 {
					zeroCol = true
				} else {
					matrix[0][col] = 0
				}
			}
		}
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
	}

	if zeroRow {
		for col := 0; col < cols; col++ {
			matrix[0][col] = 0
		}
	}

	if zeroCol {
		for row := 0; row < rows; row++ {
			matrix[row][0] = 0
		}
	}
}

// same ideas as setZeroes2
// time complexity: O(M * N)
// space complexity: O(1)
func setZeroes3(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	zeroCol := false
	for row := 0; row < rows; row++ {
		if matrix[row][0] == 0 {
			zeroCol = true
		}

		for col := 1; col < cols; col++ {
			// as a flag
			if matrix[row][col] == 0 {
				matrix[0][col] = 0
				// wether which element is zero in row[0], matrix[0][0] will be setting to zero
				// so if matrix[0][0] is zero that mean all of element in the first row must be to set zero
				matrix[row][0] = 0
			}
		}
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if matrix[0][col] == 0 || matrix[row][0] == 0 {
				matrix[row][col] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for col := 1; col < cols; col++ {
			matrix[0][col] = 0
		}
	}

	if zeroCol {
		for row := 0; row < rows; row++ {
			matrix[row][0] = 0
		}
	}
}
