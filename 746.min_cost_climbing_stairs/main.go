package main

// DP bottom-up
// time complexity: O(N)
// space complexity: O(N)
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i < len(dp); i++ {
		oneStep := dp[i-1] + cost[i-1]
		twoStep := dp[i-2] + cost[i-2]
		dp[i] = min(oneStep, twoStep)
	}

	return dp[len(dp)-1]
}

// DP top-down
// time complexity: O(N)
// space complexity: O(N)
func minCostClimbingStairs2(cost []int) int {
	mem := make(map[int]int)

	var dp func(n int) int
	dp = func(n int) int {
		if n <= 1 {
			return 0
		}

		if _, ok := mem[n]; ok {
			return mem[n]
		}

		step1 := dp(n-1) + cost[n-1]
		step2 := dp(n-2) + cost[n-2]
		mem[n] = min(step1, step2)
		return mem[n]
	}

	return dp(len(cost))
}

// Optimized DP
// time complexity: O(N)
// space complexity: O(1)
func minCostClimbingStairs3(cost []int) int {
	n1, n2 := 0, 0
	for i := 2; i <= len(cost); i++ {
		temp := n1
		n1 = min(n1+cost[i-1], n2+cost[i-2])
		n2 = temp
	}
	return n1
}
