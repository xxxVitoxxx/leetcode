package main

// time complexity: O(N*N)
// space complexity: O(1)
func longestPalindrome(s string) string {
	expend := func(left, right int) string {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		return s[left+1 : right]
	}

	var result string
	for i := range s {
		str1 := expend(i, i)
		if len(str1) > len(result) {
			result = str1
		}

		str2 := expend(i, i+1)
		if len(str2) > len(result) {
			result = str2
		}
	}

	return result
}

// DP
// time complexity: O(N*N)
// space complexity: O(N*N)
func longestPalindrome2(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	ans := []int{0, 0}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans = []int{i, i + 1}
		}
	}

	for diff := 2; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				ans = []int{i, j}
			}
		}
	}
	i, j := ans[0], ans[1]
	return s[i : j+1]
}
