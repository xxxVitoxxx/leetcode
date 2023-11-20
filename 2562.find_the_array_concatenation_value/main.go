package main

import "math"

func findTheArrayConcVal(nums []int) int64 {
	var conc int
	left, right := 0, len(nums)-1
	for left < right {
		mult, num := 1, nums[right]
		for ; num > 0; num /= 10 {
			mult *= 10
		}

		conc += nums[left]*mult + nums[right]
		left, right = left+1, right-1
	}

	if left == right {
		conc += nums[left]
	}
	return int64(conc)
}

func findTheArrayConcVal2(nums []int) int64 {
	var conc int64
	left, right := 0, len(nums)-1
	for left < right {
		mult := int(math.Log10(float64(nums[right]))) + 1
		conc += int64(nums[left])*int64(math.Pow10(mult)) + int64(nums[right])
		left, right = left+1, right-1
	}
	if left == right {
		conc += int64(nums[left])
	}
	return conc
}

// use recursion
func findTheArrayConcVal3(nums []int) int64 {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return int64(nums[0])
	}

	first, last := nums[0], nums[l-1]
	mult := int(math.Log10(float64(last))) + 1
	conc := int(math.Pow10(mult))*first + last
	return int64(conc) + findTheArrayConcVal2(nums[1:max(l-1, 1)])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
