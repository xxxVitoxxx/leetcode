package main

import (
	"golang.org/x/exp/slices"
)

// 先找到左邊比右邊大的 index 如下為 index 1(2) 和 index 2(6)
// 再從最右邊的數往回找到比 index 1(2) 還大的數並做交換，如下 index 1(2) 和 index 5(3) 交換
// 最後在針對 index 1(2) 右邊做遞增的排序
// e.g. [1, 2, 6, 5, 4, 3] -> [1, 3, 6, 5, 4, 2] -> [1, 3, 2, 4, 5, 6]
// ** 找到 index 1(2) 時，可以確定右邊會是遞減的排序，所以可以直接從最右邊往回找大於 index 1(2) 直接交換
// ** 交換後， index 1(2) 右邊仍會是遞減的排序
// ** 所以可以直接用 two point 將數組用 in-place 的方式改為遞增的排序

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			pos := i
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < nums[pos] && nums[i-1] < nums[j] {
					pos = j
				}
			}

			nums[i-1], nums[pos] = nums[pos], nums[i-1]
			slices.Sort(nums[i:])
			return
		}
	}
	slices.Sort(nums)
}

func nextPermutation2(nums []int) {
	idx := len(nums) - 1
	for ; idx > 0 && nums[idx-1] >= nums[idx]; idx-- {
	}

	if idx != 0 {
		l := len(nums) - 1
		for ; l > idx && nums[idx-1] >= nums[l]; l-- {
		}
		nums[idx-1], nums[l] = nums[l], nums[idx-1]
	}

	for r := len(nums) - 1; idx < r; idx, r = idx+1, r-1 {
		nums[idx], nums[r] = nums[r], nums[idx]
	}
}

func nextPermutation3(nums []int) {
	var found bool
	right := len(nums) - 1
	left := right - 1
	for left >= 0 {
		if !found && nums[left] < nums[right] {
			found = true
			right = len(nums) - 1
		} else if !found {
			left, right = left-1, right-1
		} else {
			if nums[left] < nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
			right--
		}
	}

	for i, j := left+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
