package main

import "strings"

// BFS
// time complexity: O(N*N*N + M*K)
// space complexity: O(N+M*K)
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	n := len(s)
	seen := make([]bool, n+1)
	queue := []int{0}

	for len(queue) > 0 {
		left := queue[0]
		queue = queue[1:]

		if left == n {
			return true
		}

		for right := left + 1; right <= n; right++ {
			if seen[right] {
				continue
			}

			if _, ok := dict[s[left:right]]; ok {
				seen[right] = true
				queue = append(queue, right)
			}
		}
	}

	return false
}

// DP bottom-up
// time complexity: O(N*M*K)
// space complexity: O(N)
func wordBreak2(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func e() {
	var str strings.Builder
	str.WriteByte(byte(2))
}
