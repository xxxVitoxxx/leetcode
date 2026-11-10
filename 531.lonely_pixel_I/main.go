package main

// time complexity: O(MN)
// space complexity: O(M+N)
func findLonelyPixel(picture [][]byte) int {
	m, n := len(picture), len(picture[0])
	seenRows, seenCols := make([]int, m), make([]int, n)

	for row := range m {
		for col := range n {
			if picture[row][col] == 'B' {
				seenRows[row]++
				seenCols[col]++
			}
		}
	}

	var count int
	for i := range seenRows {
		if seenRows[i] == 1 && seenCols[i] == 1 {
			count++
		}
	}

	return count
}

// time complexity: O(MN)
// space complexity: O(1)
func findLonelyPixel2(picture [][]byte) int {
	var count int
	m, n := len(picture), len(picture[0])
	for i := range m {
		var rowBlack, col int
		for j := range n {
			if picture[i][j] == 'B' {
				rowBlack++
				col = j
			}
		}

		if rowBlack != 1 {
			continue
		}

		var colBlack int
		for j := range m {
			if picture[j][col] == 'B' {
				colBlack++
			}
		}

		if colBlack != 1 {
			continue
		}

		count++
	}

	return count
}
