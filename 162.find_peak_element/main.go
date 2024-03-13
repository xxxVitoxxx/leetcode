package main

// linear scan
// time complexity: O(n)
// space complexity: O(1)
func findPeakElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

// recursion and binary search
// time complexity: O(logn) reduce the search space in half at every step.
// space complexity: O(logn) the depth of recursion tree will go up to logn.
func findPeakElement2(nums []int) int {
	return binarySearch(0, len(nums)-1, nums)
}

func binarySearch(left, right int, nums []int) int {
	if left == right {
		return left
	}

	mid := (left + right) / 2
	if nums[mid] > nums[mid+1] {
		return binarySearch(left, mid, nums)
	}
	return binarySearch(mid+1, right, nums)
}

// iteration and binary search
// time complexity: O(logn)
// space complexity: O(1)
func findPeakElement3(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
