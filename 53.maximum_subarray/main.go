package main

// Brute Force
// time complexity: O(N*N)
// space complexity: O(1)
func maxSubArray(nums []int) int {
	result := -10000
	for i := range len(nums) {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]

			result = max(result, sum)
		}
	}
	return result
}

// DP Kadane's Algorithm
// time complexity: O(N)
// space complexity: O(1)
func maxSubArray2(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if currSum < 0 {
			currSum = nums[i]
		} else {
			currSum += nums[i]
		}

		maxSum = max(maxSum, currSum)
	}
	return maxSum
}
