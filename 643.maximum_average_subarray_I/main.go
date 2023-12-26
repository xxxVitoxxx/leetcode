package main

import "math"

// time complexity: O(n)
// space complexity: O(n)
func findMaxAverage(nums []int, k int) float64 {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = nums[i] + prefixSum[i-1]
	}

	ans := float64(prefixSum[k-1]) / float64(k)
	for i := k; i < len(prefixSum); i++ {
		average := float64(prefixSum[i]-prefixSum[i-k]) / float64(k)
		ans = math.Max(ans, average)
	}
	return ans
}

// time complexity: O(n)
// space complexity: O(1)
func findMaxAverage2(nums []int, k int) float64 {
	var sum, ans float64 = 0, -1 << 63
	for i := 0; i < len(nums); i++ {
		sum += float64(nums[i])
		if i < k-1 {
			continue
		}
		if i >= k {
			sum -= float64(nums[i-k])
		}
		ans = math.Max(ans, sum/float64(k))
	}
	return ans
}
