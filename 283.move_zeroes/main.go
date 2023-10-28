package main

// brute force
func moveZeroes(nums []int) {
	for i := range nums {
		if nums[i] == 0 {
			j := i + 1
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j < len(nums) && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func moveZeroes2(nums []int) {
	var c int
	for i := range nums {
		if nums[i] != 0 {
			nums[c] = nums[i]
			c++
		}
	}
	for c < len(nums) {
		nums[c] = 0
		c++
	}
}

func moveZeroes3(nums []int) {
	var nonZero int
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[nonZero] = nums[nonZero], nums[i]
			nonZero++
		}
	}
}

func moveZeroes4(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
