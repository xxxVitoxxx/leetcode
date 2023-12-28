package main

// time complexity: O(n)
// space complexity: O(n)
func productExceptSelf(nums []int) []int {
	prefix, suffix := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefix[i] = 1
			continue
		}
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			suffix[i] = 1
			continue
		}
		suffix[i] = nums[i+1] * suffix[i+1]
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = prefix[i] * suffix[i]
	}
	return ans
}

// optimize productExceptSelf
// time complexity: O(n)
// space complexity: O(1)
func productExceptSelf2(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r *= nums[i]
	}
	return ans
}
