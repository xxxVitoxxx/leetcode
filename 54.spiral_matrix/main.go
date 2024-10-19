package main

import "math"

// brute force will result in a time limit exceeded error
func spiralOrder(matrix [][]int) []int {
	mixRow, mixCol := len(matrix), len(matrix[0])
	result := make([]int, mixRow*mixCol)
	row, col := 0, 0
	result[0] = matrix[row][col]
	visit := 101
	matrix[row][col] = visit
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 1; i < len(result); {
		for j := 0; j < len(direction); {
			newRow, newCol := row+direction[j][0], col+direction[j][1]
			if newRow == mixRow || newCol == mixCol ||
				newRow < 0 || newCol < 0 {
				j++
				continue
			}

			if matrix[newRow][newCol] == visit {
				j++
				continue
			}

			row, col = newRow, newCol
			result[i] = matrix[row][col]
			matrix[row][col] = visit
			i++
			if i == len(result) {
				break
			}
		}
	}

	return result
}

// optimize spiralOrder
// time complexity: O(M * N) this is because we visit each element once.
// space complexity: O(1)
//
// if we mark the cells that we visited, then when we run into a visited cell,
// we know we need to run.
//
// if `changeDirection` is larger than 1, it means we are continuously
// changing our directions, and therefore we've visited all of the elements.
//
// modifying the original data may not be a good choice sometimes. therefore,
// we can also prepare another boolean matrix to store the cells we visited.
// however, the implementation of this approach is quite similar.
func spiralOrder2(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	row, col := 0, 0
	result := make([]int, 0, m*n)
	visit := 101
	result = append(result, matrix[row][col])
	matrix[row][col] = visit

	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currentDirection := 0
	changeDirection := 0

	for changeDirection < 2 {
		for row+directions[currentDirection][0] >= 0 &&
			row+directions[currentDirection][0] < m &&
			col+directions[currentDirection][1] >= 0 &&
			col+directions[currentDirection][1] < n &&
			matrix[row+directions[currentDirection][0]][col+directions[currentDirection][1]] != visit {
			row += directions[currentDirection][0]
			col += directions[currentDirection][1]
			result = append(result, matrix[row][col])
			matrix[row][col] = visit

			// reset this to 0 since we did not break and change the direction
			changeDirection = 0
		}
		// change our direction
		currentDirection = (currentDirection + 1) % 4
		// increment changeDirection because we changed our direction
		changeDirection++
	}

	return result
}

// set up boundaries
// time complexity: O(M * N)
// space complexity: O(1)

// there are only two movement pattern, right + down and left + up.
// in the first case we increment either the row or column by 1, and
// in the latter case we increment either the row or column by -1.
// also notice that after we move horizontally the number of possible
// vertical steps decrease by 1, and after we move vertically the number
// of possible horizontal steps decrease by 1. when we run out of either
// vertical steps or horizontal steps we reach the end.
func spiralOrder3(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	direction := 1
	row, col := 0, -1
	result := []int{}
	for math.Min(float64(m), float64(n)) != 0 {
		// move horizontally
		for i := 0; i < n; i++ {
			col += direction
			result = append(result, matrix[row][col])
		}
		m -= 1

		// move vertically
		for i := 0; i < m; i++ {
			row += direction
			result = append(result, matrix[row][col])
		}
		n -= 1

		// flip direction
		direction *= -1
	}
	return result
}
