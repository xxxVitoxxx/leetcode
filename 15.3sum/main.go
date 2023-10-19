package main

import (
	"reflect"
	"sort"
)

// brute force
func threeSum(nums []int) [][]int {
	var res [][]int
	n1, n2 := 0, 1
	for n3 := 2; n3 < len(nums); n3++ {
		if nums[n1]+nums[n2]+nums[n3] == 0 {
			min := min(nums[n1], nums[n2], nums[n3])
			middle := middle(nums[n1], nums[n2], nums[n3])
			max := max(nums[n1], nums[n2], nums[n3])
			if !Exist(res, []int{min, middle, max}) {
				res = append(res, []int{min, middle, max})
			}
		}

		if n3 == len(nums)-1 {
			if n2 == len(nums)-2 {
				n1++
				n2 = n1 + 1
			} else {
				n2++
			}
			n3 = n2
		}
	}
	return res
}

func min(n1, n2, n3 int) int {
	if n1 <= n2 && n1 <= n3 {
		return n1
	} else if n2 <= n1 && n2 <= n3 {
		return n2
	}
	return n3
}

func middle(n1, n2, n3 int) int {
	if n1 <= n2 && n1 >= n3 || n1 >= n2 && n1 <= n3 {
		return n1
	} else if n2 <= n1 && n2 >= n3 || n2 >= n1 && n2 <= n3 {
		return n2
	}
	return n3
}

func max(n1, n2, n3 int) int {
	if n1 >= n2 && n1 >= n3 {
		return n1
	} else if n2 >= n1 && n2 >= n3 {
		return n2
	}
	return n3
}

func Exist(res [][]int, nums []int) bool {
	var exist bool
	for i := range res {
		if reflect.DeepEqual(res[i], nums) {
			exist = true
		}
	}
	return exist
}

// brute force
func threeSum2(nums []int) [][]int {
	tripletMap := make(map[[3]int][]int)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					current := [3]int{nums[i], nums[j], nums[k]}
					sort.Ints(current[:])
					tripletMap[current] = current[:]
				}
			}
		}
	}

	var res [][]int
	for _, triplet := range tripletMap {
		res = append(res, triplet)
	}
	return res
}

// use two pointers
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first-1] == nums[first] {
			continue
		}

		second, third := first+1, len(nums)-1
		for second < third {
			current := nums[first] + nums[second] + nums[third]
			if current < 0 {
				second++
			} else if current > 0 {
				third--
			} else {
				res = append(res, []int{nums[first], nums[second], nums[third]})
				second++
				third--

				for second < third && nums[second-1] == nums[second] {
					second++
				}

				for second < third && nums[third] == nums[third-1] {
					third--
				}
			}
		}
	}
	return res
}
