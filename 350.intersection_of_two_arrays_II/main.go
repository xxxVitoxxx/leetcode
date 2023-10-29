package main

import "sort"

// brute force
func intersect(nums1, nums2 []int) []int {
	m1 := make(map[int]int)
	for _, num := range nums1 {
		m1[num]++
	}

	m2 := make(map[int]int)
	for _, num := range nums2 {
		m2[num]++
	}

	var res []int
	for k, v := range m1 {
		if nums2Val, ok := m2[k]; ok {
			for c := min(v, nums2Val); c > 0; c-- {
				res = append(res, k)
			}
		}
	}
	return res
}

func min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	}
	return n2
}

// use two pointers
func intersect2(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var res []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i, j = i+1, j+1
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}

// use in-place pattern
func intersect3(nums1, nums2 []int) []int {
	var index int
	for i := 0; i < len(nums1); i++ {
		for j := index; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				nums2[index], nums2[j] = nums2[j], nums2[index]
				index++
				break
			}
		}
	}
	return nums2[:index]
}

func intersect4(nums1, nums2 []int) []int {
	var res []int
loop:
	for _, num1 := range nums1 {
		for i, num2 := range nums2 {
			if num1 == num2 {
				res = append(res, num1)
				if len(nums2) > 0 {
					nums2 = append(nums2[:i], nums2[i+1:]...)
				} else {
					break loop
				}
				break
			}
		}
	}
	return res
}
