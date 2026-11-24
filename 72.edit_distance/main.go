package main

// Recursion
// time complexity: O(MN)
// space complexity: O(MN)
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	memo := make([][]int, n1+1)
	for i := range memo {
		memo[i] = make([]int, n2+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return j
		}

		if j == 0 {
			return i
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if word1[i-1] == word2[j-1] {
			memo[i][j] = dfs(i-1, j-1)
		} else {
			memo[i][j] = min(
				dfs(i, j-1),   // insert
				dfs(i-1, j),   // delete
				dfs(i-1, j-1), // replace
			) + 1
		}

		return memo[i][j]
	}

	return dfs(n1, n2)
}

// DP
// time complexity: O(MN)
// space complexity: O(MN)
func minDistance2(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j],
					dp[i-1][j-1],
					dp[i][j-1],
				) + 1
			}
		}
	}

	return dp[n1][n2]
}
