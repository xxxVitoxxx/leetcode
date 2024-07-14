package main

// dynamic programming (top down)
// time complexity: O(n). we recursively call `dfs` on subproblem and
// each subproblem `dfs(i)` is computed once.
//
// space complexity: O(n). the hash map `dp` contains at most `n+1` key-value pairs.
func tribonacci(n int) int {
	dp := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}
	var dfs func(int) int
	dfs = func(n int) int {
		if _, ok := dp[n]; ok {
			return dp[n]
		}
		dp[n] = dfs(n-1) + dfs(n-2) + dfs(n-3)
		return dp[n]
	}
	return dfs(n)
}
