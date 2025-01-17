package main

// time complexity: O(n^2)
// space complexity: O(n)
func getRow(rowIndex int) []int {
	prev := []int{1}
	for row := 1; row <= rowIndex; row++ {
		curr := make([]int, row+1)
		curr[0], curr[row] = 1, 1
		for col := 1; col < row; col++ {
			curr[col] = prev[col-1] + prev[col]
		}
		prev = curr
	}
	return prev
}

type key struct {
	row int
	col int
}

var cache = make(map[key]int)

func getNum(row, col int) int {
	rowCol := key{row, col}
	if val, ok := cache[rowCol]; ok {
		return val
	}
	if row == 0 || col == 0 || col == row {
		cache[rowCol] = 1
		return cache[rowCol]
	}
	cache[rowCol] = getNum(row-1, col-1) + getNum(row-1, col)
	return cache[rowCol]
}

// time complexity: O(n^2)
// space complexity: O(n)
func getRow2(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	for i := range ans {
		ans[i] = getNum(rowIndex, i)
	}
	return ans
}
