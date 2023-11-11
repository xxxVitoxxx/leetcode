package main

import "sort"

func sortEvenOdd(nums []int) []int {
	evenIdx, oddIdx := []int{}, []int{}
	for i, num := range nums {
		if i%2 == 0 {
			evenIdx = append(evenIdx, num)
		} else {
			oddIdx = append(oddIdx, num)
		}
	}

	sort.Ints(evenIdx)
	sort.Slice(oddIdx, func(i, j int) bool {
		return oddIdx[i] > oddIdx[j]
	})

	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		if i%2 == 0 {
			res[i], evenIdx = evenIdx[0], evenIdx[1:]
		} else {
			res[i], oddIdx = oddIdx[0], oddIdx[1:]
		}
	}
	return res
}

func sortEvenOdd2(nums []int) []int {
	// [4,1,2,3]
	for i := 1; i < len(nums); i += 2 {
		nums[i] = -nums[i]
	}
	// [4,-1,2,-3]

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// [-3, -1, 2, 4]

	res := make([]int, len(nums))
	even, odd := 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			res[odd] = -nums[i]
			odd += 2
		} else {
			res[even] = nums[i]
			even += 2
		}
	}
	return res
}

func sortEvenOdd3(nums []int) []int {
	// sort even index
	for i := 0; i < len(nums); i += 2 {
		minIdx := i
		for j := i + 2; j < len(nums); j += 2 {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		if i != minIdx {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}

	// sort odd index
	for i := 1; i < len(nums); i += 2 {
		maxIdx := i
		for j := i + 2; j < len(nums); j += 2 {
			if nums[j] > nums[maxIdx] {
				maxIdx = j
			}
		}
		if i != maxIdx {
			nums[i], nums[maxIdx] = nums[maxIdx], nums[i]
		}
	}
	return nums
}

func sortEvenOdd4(nums []int) []int {
	for i := 0; i < len(nums)-1; i += 2 {
		for j := i + 2; j < len(nums); j += 2 {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}

			if j+1 == len(nums) {
				continue
			}

			if nums[i+1] < nums[j+1] {
				nums[i+1], nums[j+1] = nums[j+1], nums[i+1]
			}
		}
	}
	return nums
}
