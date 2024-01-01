package main

// sliding window
// time complexity: O(n)
// space complexity: O(1)
func longestSubarray(nums []int) int {
	ans, zero := 0, 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		}
		for zero > 1 {
			if nums[left] == 0 {
				zero--
			}
			left++
		}

		ans = max(ans, right-left)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// sliding window
// time complexity: O(n)
// space complexity: O(1)
func longestSubarray2(nums []int) int {
	left, right, zero := 0, 0, 0
	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		}
		if zero > 1 {
			if nums[left] == 0 {
				zero--
			}
			left++
		}
	}
	return right - left - 1
}
