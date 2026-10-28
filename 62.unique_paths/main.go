package main

// this solution will timeout
// time complexity: O(2^(M+N))
// space complexity: O(M+N)
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

// use memoization
func uniquePaths2(m int, n int) int {
	return helper(m, n, make(map[[2]int]int))
}

func helper(m, n int, memo map[[2]int]int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if value, ok := memo[[2]int{m, n}]; ok {
		return value
	}

	memo[[2]int{m - 1, n}] = helper(m-1, n, memo)
	memo[[2]int{m, n - 1}] = helper(m, n-1, memo)
	return memo[[2]int{m - 1, n}] + memo[[2]int{m, n - 1}]
}

// DP
// time complexity: O(MN)
// space complexity: O(MN)
func uniquePaths3(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}

	return dp[m-1][n-1]
}
