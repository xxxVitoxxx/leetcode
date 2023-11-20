package main

// brute force
func sortColors(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		pos := i
		for j := i + 1; j < len(nums); j++ {
			if nums[pos] > nums[j] {
				pos = j
			}
		}
		if pos != i {
			nums[i], nums[pos] = nums[pos], nums[i]
		}
	}
}

// use two pointer
func sortColors2(nums []int) {
	left, right := 0, len(nums)-1
	var idx int
	for idx <= right {
		if nums[idx] == 0 {
			nums[left], nums[idx] = nums[idx], nums[left]
			left, idx = left+1, idx+1
		} else if nums[idx] == 2 {
			nums[right], nums[idx] = nums[idx], nums[right]
			right--
		} else {
			idx++
		}
	}
}
