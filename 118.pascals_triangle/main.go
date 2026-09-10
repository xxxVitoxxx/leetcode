package main

/*
1
1, 1
1, 2, 1
1, 3, 3, 1
1, 4, 6, 4, 1,
1, 5, 10, 10, 5, 1,

time complexity: O(N*N)
space complexity: O(N*N)
*/
func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	triangle[0] = append(triangle[0], 1)

	for row := 1; row < numRows; row++ {
		r := make([]int, 0, row+1)
		preRow := triangle[row-1]

		r = append(r, 1)
		for i := 1; i < row; i++ {
			r = append(r, preRow[i-1]+preRow[i])
		}

		r = append(r, 1)
		triangle[row] = r
	}
	return triangle
}

// recursion
func generate2(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	triangle := generate2(numRows - 1)
	preRow := triangle[len(triangle)-1]

	currRow := make([]int, 0, len(triangle))
	currRow = append(currRow, 1)
	for i := 1; i < len(triangle); i++ {
		currRow = append(currRow, preRow[i-1]+preRow[i])
	}

	currRow = append(currRow, 1)
	triangle = append(triangle, currRow)

	return triangle
}
