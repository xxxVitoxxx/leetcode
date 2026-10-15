package main

// time complexity: O(N*N*N)
// space complexity: O(1)
func countSubstrings(s string) int {
	var result int
	for left := 0; left < len(s); left++ {
		for right := left; right < len(s); right++ {
			if isPalindromic(s, left, right) {
				result++
			}
		}
	}

	return result
}

func isPalindromic(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// DP
// time complexity: O(N*N)
// space complexity: O(N*N)
func countSubstrings2(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	var count int
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}

// expand around possible center
// time complexity: O(N*N)
// space complexity: O(1)
func countSubstrings3(s string) int {
	var result int
	for i := range s {
		result += expand(s, i, i)
		result += expand(s, i, i+1)
	}
	return result
}

func expand(s string, left, right int) int {
	var result int
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}

		left--
		right++
		result++
	}

	return result
}
