package main

func removeElement(nums []int, val int) int {
	var res int
	for _, num := range nums {
		if num != val {
			nums[res] = num
			res++
		}
	}
	return res
}

// use two pointer
func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}

		for left <= right && nums[right] == val {
			right--
		}

		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}
	return left
}
