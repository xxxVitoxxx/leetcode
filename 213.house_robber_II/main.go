package main

// DP
// time complexity: O(N)
// space complexity: O(N)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(doRob(nums[:len(nums)-1]), doRob(nums[1:]))
}

func doRob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// DP
// time complexity: O(N)
// space complexity: O(1)
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(doRob2(nums[:len(nums)-1]), doRob2(nums[1:]))
}

func doRob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		first, second = second, max(first+nums[i], second)
	}

	return second
}
