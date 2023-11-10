package main

func sortArrayByParity(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(res)-1
	for _, num := range nums {
		if left <= right {
			if num%2 == 0 {
				res[left] = num
				left++
			} else {
				res[right] = num
				right--
			}
		}
	}
	return res
}

func sortArrayByParity2(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]%2 == 0 {
			left++
		}
		for left < right && nums[right]%2 == 1 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}

func sortArrayByParity3(nums []int) []int {
	var evenIdx int
	for i := range nums {
		if nums[i]%2 == 0 {
			nums[i], nums[evenIdx] = nums[evenIdx], nums[i]
			evenIdx++
		}
	}
	return nums
}
