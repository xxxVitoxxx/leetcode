package main

// DFS
// time complexity: O(N)
// space complexity: O(N)
func numDecodings(s string) int {
	n := len(s)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i == n {
			return 1
		}

		if s[i] == '0' {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		result := dfs(i + 1)
		if i+1 < n {
			twoDigit := int(s[i]-'0')*10 + int(s[i]-'0')
			if twoDigit >= 10 && twoDigit <= 26 {
				result += dfs(i + 2)
			}
		}

		memo[i] = result

		return result
	}

	return dfs(0)
}

// DP
// time complexity: O(N)
// space complexity: O(N)
func numDecodings2(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		if i >= 2 {
			twoDigit := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
			if twoDigit >= 10 && twoDigit <= 26 {
				dp[i] += dp[i-2]
			}
		}

	}
	return dp[n]
}
