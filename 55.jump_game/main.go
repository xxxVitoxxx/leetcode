package main

// Greedy
func canJump(nums []int) bool {
	var farthest int
	for i := range nums {
		if i > farthest {
			return false
		}

		farthest = max(farthest, i+nums[i])
	}
	return true
}

// Greedy
func canJump2(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
