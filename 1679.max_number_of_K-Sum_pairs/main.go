package main

import "golang.org/x/exp/slices"

// use two pointer
// time complexity: O(nlogn)
// space complexity: O(1)
func maxOperations(nums []int, k int) int {
	slices.Sort(nums)
	ans, left, right := 0, 0, len(nums)-1
	for left < right {
		num := nums[left] + nums[right]
		if num == k {
			ans++
			left, right = left+1, right-1
		} else if num < k {
			left++
		} else {
			right--
		}
	}
	return ans
}

// use hash
// time complexity: O(n)
// space complexity: O(n)
func maxOperations2(nums []int, k int) int {
	ans, m := 0, make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[k-nums[i]] > 0 {
			ans++
			m[k-nums[i]]--
		} else {
			m[nums[i]]++
		}
	}
	return ans
}
