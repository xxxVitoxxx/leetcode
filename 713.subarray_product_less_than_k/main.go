package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans, prod := 0, 1
	for left, right := 0, 0; right < len(nums); right++ {
		prod *= nums[right]
		for left <= right && prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
