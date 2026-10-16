package main

// Brute Force
// time complexity: O(N*N)
// space complexity: O(1)
func maxProduct(nums []int) int {
	result := nums[0]
	for range len(nums) {
		n := 1
		for _, num := range nums {
			n *= num
			result = max(result, n)
		}
	}
	return result
}

// DP
// time complexity: O(N)
// space complexity: O(1)
func maxProduct2(nums []int) int {
	currMin, currMax, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		n1, n2 := currMin*nums[i], currMax*nums[i]

		currMin = min(nums[i], n1, n2)
		currMax = max(nums[i], n1, n2)

		result = max(result, currMax)
	}
	return result
}
