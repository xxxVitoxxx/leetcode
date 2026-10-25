package main

// DP
// time complexity: O(N*N)
// space complexity: O(N)
func lengthOfLIS(nums []int) int {
	var result int
	dp := make([]int, len(nums))
	for i := range len(nums) {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		result = max(result, dp[i])
	}
	return result
}

// Binary Search
// time complexity: O(N log N)
// space complexity: O(N)
func lengthOfLIS2(nums []int) int {
	tails := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		index := binarySearch(tails, nums[i])
		if index == len(tails) {
			tails = append(tails, nums[i])
		} else {
			tails[index] = nums[i]
		}
	}
	return len(tails)
}

func binarySearch(tails []int, target int) int {
	left, right := 0, len(tails)
	for left < right {
		mid := left + (right-left)/2

		if tails[mid] < target {
			left = mid + 1
		} else if tails[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
}
