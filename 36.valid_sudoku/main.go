package main

// time complexity: O(N*N)
// space complexity: O(N*N)
func isValidSudoku(board [][]byte) bool {
	n := len(board)
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	boxes := make([]map[byte]bool, n)
	for i := range n {
		rows[i] = make(map[byte]bool, n)
		cols[i] = make(map[byte]bool, n)
		boxes[i] = make(map[byte]bool, n)
	}

	for row := range n {
		for col := range n {
			value := board[row][col]
			if value == '.' {
				continue
			}

			if rows[row][value] {
				return false
			}
			rows[row][value] = true

			if cols[col][value] {
				return false
			}
			cols[col][value] = true

			boxKey := getBoxKey(row, col, n/3)
			if boxes[boxKey][value] {
				return false
			}
			boxes[boxKey][value] = true
		}
	}

	return true
}

func getBoxKey(row, col, n int) int {
	return row/n*n + col/n
}
