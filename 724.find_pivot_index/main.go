package main

// prefix sum
// time complexity: O(n)
// space complexity: O(1)
func pivotIndex(nums []int) int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	l := len(prefix) - 1
	for i := 0; i < len(prefix); i++ {
		if i == 0 && prefix[l]-prefix[i] == 0 {
			return i
		}
		if i > 0 && prefix[i-1] == prefix[l]-prefix[i] {
			return i
		}
	}
	return -1
}

// optimize pivotIndex
// prefix sum
// time complexity: O(n)
// space complexity: O(1)
func pivotIndex2(nums []int) int {
	sum, leftSum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
