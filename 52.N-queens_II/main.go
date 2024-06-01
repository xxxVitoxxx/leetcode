package main

// we should never place a queen on an attacked square.
// when all possibilities are exhausted, backtracking by removing a queen and
// placing it elsewhere.

// a queen can be attacked if another queen is in any of the 4 following positions:
// on the same row, on the same columns, on the same diagonal, or on the same anti-diagonal.
// to implement backtracking, we implement a backtrack function that makes some changes to the state,
// calls itself again, and then when that call returns it undoes those changes.

// each time our backtrack function is called, we can encode state in the following manner:
//   - to make sure that we only place 1 queen per row, we will pass an integer argument row in to backtrack,
//     and will only place one queen during each call. whenever we place a queen, we will move onto the next row
//     by calling backtrack again with the parameter value row+1.
//   - to make sure we only place 1 queen per column, we will use a set. whenever we place a queen,
//     we can add the column index to this set.

// the diagonals are a little trickier, but they have a property that we can use to our advantage.
// for each square on given diagonal, the difference between the row and column indexes (row-col) will be constant.
// think about the diagonal that starts from (0, 0). the i square has coordinates (i, i), so the difference is always 0.
//     0   1   2   3   4
//    -------------------
//  0| 0 |-1 |-2 |-3 |-4 |
//    -------------------
//  1| 1 | 0 |-1 |-2 |-3 |
//    -------------------
//  2| 2 | 1 | 0 |-1 |-2 |
//    -------------------
//  3| 3 | 2 | 1 | 0 |-1 |
//    -------------------
//  4| 4 | 3 | 2 | 1 | 0 |
//    -------------------
//
// for each square on a given anti-diagonal, the sum of the row and column indexes (row+col) will be constant.
// if you were to start at the highest square in an anti-diagonal and move downwards,
// the row index increments by 1 (row+1), and the column index decrements by 1 (col-1).
// these cancel each other out.
//     0   1   2   3   4
//    -------------------
//  0| 0 | 1 | 2 | 3 | 4 |
//    -------------------
//  1| 1 | 2 | 3 | 4 | 5 |
//    -------------------
//  2| 2 | 3 | 4 | 5 | 6 |
//    -------------------
//  3| 3 | 4 | 5 | 6 | 7 |
//    -------------------
//  4| 4 | 5 | 6 | 7 | 8 |
//    -------------------
//
// whenever we place a queen, we should calculate the diagonal and the anti-diagonal value it belongs to.
// in the same way we had a set for the column, we should also have a set for both the diagonals and anti-diagonals.
// then, we can put the values for this square into the corresponding sets.

// time complexity: O(n!), where n is the number of queens (which is the same as the width and height of the board)
// space complexity: O(n)
// extra memory used includes the 3 sets used to store board state,
// as well as the recursion call stack. all of this scales linearly with the number of queens.
func totalNQueens(n int) int {
	var backtrack func(int, map[int]bool, map[int]bool, map[int]bool) int
	backtrack = func(row int, diagonals, antiDiagonals, cols map[int]bool) int {
		if row == n {
			return 1
		}

		var solution int
		for col := 0; col < n; col++ {
			currDiagonal := row - col
			currAntiDiagonal := row + col
			if cols[col] || diagonals[currDiagonal] || antiDiagonals[currAntiDiagonal] {
				continue
			}

			cols[col] = true
			diagonals[currDiagonal] = true
			antiDiagonals[currAntiDiagonal] = true
			solution += backtrack(row+1, diagonals, antiDiagonals, cols)

			delete(cols, col)
			delete(diagonals, currDiagonal)
			delete(antiDiagonals, currAntiDiagonal)
		}

		return solution
	}

	return backtrack(
		0,
		make(map[int]bool),
		make(map[int]bool),
		make(map[int]bool),
	)
}

func totalNQueens2(n int) int {
	var backtrack func(int, int, int, int) int
	backtrack = func(row, cols, diagonals, antiDiagonals int) int {
		if row == n {
			return 1
		}

		var solution int
		for col := 0; col < n; col++ {
			currCol := 1 << col
			currDiagonal := 1 << (row - col + n)
			currAntiDiagonal := 1 << (row + col)
			if cols&currCol > 0 ||
				diagonals&currDiagonal > 0 ||
				antiDiagonals&currAntiDiagonal > 0 {
				continue
			}

			cols ^= currCol
			diagonals ^= currDiagonal
			antiDiagonals ^= currAntiDiagonal
			solution += backtrack(row+1, cols, diagonals, antiDiagonals)

			cols ^= currCol
			diagonals ^= currDiagonal
			antiDiagonals ^= currAntiDiagonal
		}

		return solution
	}

	return backtrack(0, 0, 0, 0)
}
