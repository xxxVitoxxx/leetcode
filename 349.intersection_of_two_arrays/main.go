package main

import (
	"sort"

	"golang.org/x/exp/slices"
)

// brute force
func intersection(nums1, nums2 []int) []int {
	var res []int
	m := make(map[int]bool)
	for i := range nums1 {
		if m[nums1[i]] {
			continue
		}
		m[nums1[i]] = true
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				res = append(res, nums2[j])
				break
			}
		}
	}
	return res
}

func intersection2(nums1, nums2 []int) []int {
	m := make(map[int]bool)
	for _, n := range nums1 {
		if !m[n] {
			m[n] = true
		}
	}

	var res []int
	for _, n := range nums2 {
		if m[n] && !slices.Contains(res, n) {
			res = append(res, n)
		}
	}
	return res
}

func intersection3(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var res []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			num := nums1[i]
			res = append(res, num)
			i, j = i+1, j+1

			for i < len(nums1) && nums1[i] == num {
				i++
			}

			for j < len(nums2) && nums2[j] == num {
				j++
			}
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}

func intersection4(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var res []int
	last := -1
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			if nums1[i] != last {
				res = append(res, nums1[i])
				last = nums1[i]
			}
			i, j = i+1, j+1
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}
