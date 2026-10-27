package main

// Greedy
// time complexity :O(N)
// space complexity: O(1)
func jump(nums []int) int {
	var result, currEnd, currFar int
	for i := range len(nums) - 1 {
		currFar = max(currFar, i+nums[i])

		if i == currEnd {
			result++
			currEnd = currFar
		}
	}
	return result
}
