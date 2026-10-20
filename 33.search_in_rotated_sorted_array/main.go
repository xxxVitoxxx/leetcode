package main

// Binary Search
// time complexity: O(log N)
// space complexity: O(1)
func search(nums []int, target int) int {
	n := len(nums) - 1
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] >= target && nums[n] >= target {
		right = n
	} else {
		left, right = 0, left
	}

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 13
		}
	}

	return -1
}

// Binary Search
// time complexity: O(log N)
// space complexity: O(1)
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// left has sort
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			// right has sort
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return -1
}
