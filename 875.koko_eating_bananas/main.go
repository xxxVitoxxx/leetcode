package main

import "math"

// binary search
// time complexity: O(n * logm) where n is length of piles, and m is max(piles)
// space complexity: O(1)
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, maxElement(piles)
	for left < right {
		hour := 0
		mid := (left + right) / 2
		for _, pile := range piles {
			hour += int(math.Ceil(float64(pile) / float64(mid)))
		}

		if hour <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func maxElement(piles []int) int {
	var max int
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	return max
}
