package main

// DP bottom-up
// M is the amount, N is the number of coin denominations
// time complexity: O(MN)
// space complexity: O(M)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for _, coin := range coins {
		for a := 1; a <= amount; a++ {
			if a >= coin {
				dp[a] = min(dp[a], dp[a-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// DP top-down
// M is the amount, N is the number of coin denominations
// time complexity: O(MN)
// space complexity: O(M)
func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount+1)

	var dfs func(int) int
	dfs = func(remaining int) int {
		if remaining == 0 {
			return 0
		}

		if memo[remaining] != 0 {
			return memo[remaining]
		}

		best := amount + 1
		for _, coin := range coins {
			if remaining >= coin {
				sub := dfs(remaining - coin)
				if sub != -1 {
					best = min(best, sub+1)
				}
			}
		}

		// 湊不出來
		if best == amount+1 {
			best = -1
		}

		memo[remaining] = best
		return best
	}

	return dfs(amount)
}
