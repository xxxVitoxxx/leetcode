package main

// DP (bottom-up)
// time complexity: O(N)
// space complexity: O(N)
func rob(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// DP (top-down)
// time complexity: O(N)
// space complexity: O(N)
func rob2(nums []int) int {
	mem := make(map[int]int)

	var dfs func(n int) int
	dfs = func(n int) int {
		if n == 0 {
			return nums[0]
		}
		if n == 1 {
			return max(nums[1], nums[0])
		}

		if _, ok := mem[n]; !ok {
			mem[n] = max(dfs(n-1), dfs(n-2)+nums[n])
		}

		return mem[n]
	}

	dfs(len(nums) - 1)

	return dfs(len(nums) - 1)
}

// Optimized DP
// time complexity: O(N)
// space complexity: O(1)
func rob3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	n1, n2 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		num := max(n1+nums[i], n2)
		n1, n2 = n2, num
	}

	return n2
}
