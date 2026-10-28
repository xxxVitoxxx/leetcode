package main

// Backtracking
// time complexity: O(M*N*3^L). L is the length of the word to be matched.
// space complexity: O(L)
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var backtracking func(int, int, int) bool
	backtracking = func(row, col, index int) bool {
		if row < 0 || col < 0 || row >= m || col >= n || index >= len(word) || board[row][col] != word[index] {
			return false
		}

		if index == len(word)-1 {
			return true
		}

		temp := board[row][col]
		board[row][col] = '$'

		nextIndex := index + 1
		found := backtracking(row+1, col, nextIndex) || backtracking(row-1, col, nextIndex) || backtracking(row, col+1, nextIndex) || backtracking(row, col-1, nextIndex)

		board[row][col] = temp

		return found
	}

	for row := range m {
		for col := range n {
			if board[row][col] != word[0] {
				continue
			}

			if backtracking(row, col, 0) {
				return true
			}
		}
	}

	return false
}
