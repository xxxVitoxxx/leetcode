package main

// N is amount, C is the length of coins
// time complexity: O(NC)
// space complexity: O(N)
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			dp[a] += dp[a-coin]
		}
	}

	return dp[amount]
}
