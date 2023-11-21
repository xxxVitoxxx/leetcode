package main

func arithmeticTriplets(nums []int, diff int) int {
	var res int
	for i := 1; i < len(nums)-1; i++ {
		left, right := 0, len(nums)-1
		for left < i && i < right {
			if nums[i]-nums[left] > diff {
				left++
			} else if nums[right]-nums[i] > diff {
				right--
			} else if nums[i]-nums[left] == diff && nums[right]-nums[i] == diff {
				res++
				break
			} else if nums[i]-nums[left] < diff || nums[right]-nums[i] < diff {
				break
			}
		}
	}
	return res
}

func arithmeticTriplets2(nums []int, diff int) int {
	res, l := 0, len(nums)
	for i, j, k := 0, 0, 0; i < l && j < l && k < l; i++ {
		target1 := nums[i] + diff
		target2 := nums[i] + 2*diff
		for j < l && nums[j] < target1 {
			j++
		}
		if j < l && nums[j] == target1 {
			for k < l && nums[k] < target2 {
				k++
			}
			if k < l && nums[k] == target2 {
				res++
				continue
			}
		}
	}
	return res
}

func arithmeticTriplets3(nums []int, diff int) int {
	var res int
	for first := 0; first < len(nums)-2; first++ {
		target1 := nums[first] + diff
		target2 := target1 + diff
		second := first + 1
		for second < len(nums)-1 && nums[second] < target1 {
			second++
		}

		third := second + 1
		for third < len(nums)-1 && nums[third] < target2 {
			third++
		}

		if second >= len(nums) || third >= len(nums) {
			continue
		}

		if nums[second] == target1 && nums[third] == target2 {
			res++
		}
	}
	return res
}

// use hash table
func arithmeticTriplets4(nums []int, diff int) int {
	res, m := 0, make(map[int]bool, len(nums))
	for _, num := range nums {
		if m[num-diff] && m[num-2*diff] {
			res++
		}
		m[num] = true
	}
	return res
}
