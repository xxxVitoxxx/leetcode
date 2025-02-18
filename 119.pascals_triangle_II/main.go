package main

// time complexity: O(n^2)
// space complexity: O(n)
//
// it is worth noting that generating a number for a particular row
// requires only two numbers from the previous row.
// consequently, generating a row only requires numbers from the previous row.
// thus, we could reduce our memory footprint by only keeping
// the latest row generated, and use that to generate a new row.
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

// dynamic programming + memoization
// time complexity: O(n^2)
// space complexity: O(n)
//
// store the intermediate results in the cache
// so that we could reuse them later without re-calculation
// to provides significant savings on runtime
func getRow2(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	for i := range ans {
		ans[i] = getNum(rowIndex, i)
	}
	return ans
}

// memory-efficient dynamic programming
// time complexity: O(n^2)
// space complexity: O(n)
// no extra space is used other than required to hold the output
//
// this is a more efficient version of the previous solution
// it uses a single array to store the intermediate results
// and updates it in place
// although there is no savings in theoretical computational complexity,
// in practice there are some minor wins:
//   - we have one vector/array instead of two. so memory consumption is roughly half.
//   - no time wasted in swapping references to vectors for previous and current row.
//   - locality of reference shines through here. since every read is for consecutive
//     memory locations in the vector/array, we get a performance boost.
func getRow3(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		ans[i] = 1
	}
	for i := 1; i < rowIndex; i++ {
		for j := i; j > 0; j-- {
			ans[j] += ans[j-1]
		}
	}
	return ans
}
