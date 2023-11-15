package main

import (
	"sort"
)

// brute force
func findTheDistanceValue(arr1, arr2 []int, d int) int {
	var res int
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				break
			}
			if j == len(arr2)-1 {
				res++
			}
		}
	}
	return res
}

// sort and binary search
func findTheDistanceValue2(arr1, arr2 []int, d int) int {
	sort.Ints(arr2)
	var res int
	for _, num := range arr1 {
		left, right := 0, len(arr2)-1
		for left <= right {
			mid := (left + right) / 2
			if abs(num-arr2[mid]) <= d {
				break
			} else if arr2[mid] < num {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left > right {
			res++
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 先用 arr1 中的數字 num 加減 d 得到 low 和 high
// 如果 arr2 中有數字大於等於 low 和小於等於 high
// 則代表 num - arr2 中的數字的絕對值會小於等於 d
// sort and binary search
func findTheDistanceValue3(arr1, arr2 []int, d int) int {
	sort.Ints(arr2)
	var res int
	for _, num := range arr1 {
		low, high := num-d, num+d
		if !binarySearch(low, high, arr2) {
			res++
		}
	}
	return res
}

func binarySearch(low, high int, arr []int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		// mid := (left + right) >> 1
		mid := (left + right) / 2
		if arr[mid] >= low && arr[mid] <= high {
			return true
		} else if arr[mid] < low {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
