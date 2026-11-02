package main

// DP
// M is the length of text1. N is the length of text2.
// time complexity: O(MN)
// space complexity: O(MN)
//
/*
   "abcde" vs "ace"

         ""   a   c   e
   ""     0   0   0   0
   a      0  [1]  1   1
   b      0   1   1   1
   c      0   1  [2]  2
   d      0   1   2   2
   e      0   1   2  [3]
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
