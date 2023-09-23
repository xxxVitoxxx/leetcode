package main

func nextGreaterElement(nums1, nums2 []int) []int {
	m := make(map[int]int, len(nums2))
	for i, v := range nums2 {
		for j := i + 1; j < len(nums2); j++ {
			if v < nums2[j] {
				m[v] = nums2[j]
				break
			}
		}
		if m[v] == 0 {
			m[v] = -1
		}
	}

	res := make([]int, len(nums1))
	for i, v := range nums1 {
		val := m[v]
		res[i] = val
	}
	return res
}
