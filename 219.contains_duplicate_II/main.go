package main

func containsNearByDuplicate(nums []int, k int) bool {
	for i := 0; i < k; i++ {
		left, right := i, i+1
		for left < len(nums) && right < len(nums) && right < left+k+1 {
			if nums[left] == nums[right] {
				return true
			}
			right++
			if right == left+k+1 {
				left = right - 1
			}
		}
	}
	return false
}

func containsNearByDuplicate2(nums []int, k int) bool {
	for pnt1 := 0; pnt1 < len(nums); pnt1++ {
		for pnt2 := pnt1; pnt2 <= pnt1+k && pnt2 < len(nums); pnt2++ {
			if pnt1 != pnt2 && nums[pnt1] == nums[pnt2] {
				return true
			}
		}
	}
	return false
}

func containsNearByDuplicate3(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, ok := m[nums[i]]; ok && i-idx <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
