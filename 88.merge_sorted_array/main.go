package main

import "sort"

// brute force
func merge(nums1 []int, m int, nums2 []int, n int) {
	for _, num2 := range nums2 {
		nums1[m] = num2
		m++
	}
	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	mi, ni := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if mi < 0 {
			for ; i >= 0; i-- {
				nums1[i] = nums2[ni]
				ni--
			}
			break
		}

		if ni < 0 {
			for ; i >= 0; i-- {
				nums1[i] = nums1[mi]
				mi--
			}
			break
		}

		if nums1[mi] >= nums2[ni] {
			nums1[i] = nums1[mi]
			mi--
		} else if nums1[mi] < nums2[ni] {
			nums1[i] = nums2[ni]
			ni--
		}
	}
}

func merge3(nums1 []int, m int, nums2 []int, n int) {
	i, mi, ni := m+n-1, m-1, n-1
	for mi >= 0 && ni >= 0 {
		if nums1[mi] >= nums2[ni] {
			nums1[i] = nums1[mi]
			mi--
			i--
		} else if nums1[mi] < nums2[ni] {
			nums1[i] = nums2[ni]
			ni--
			i--
		}
	}

	for mi >= 0 {
		nums1[i] = nums1[mi]
		mi--
		i--
	}

	for ni >= 0 {
		nums1[i] = nums2[ni]
		ni--
		i--
	}
}

func merge4(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
