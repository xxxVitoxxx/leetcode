package main

// hash
// time complexity: O(n + m)
// space complexity: O(max(n, m))
func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map, nums2Map := unique(nums1), unique(nums2)
	return [][]int{inOneNotInTwo(nums1Map, nums2Map), inOneNotInTwo(nums2Map, nums1Map)}
}

func unique(nums []int) map[int]bool {
	uniq := make(map[int]bool, len(nums))
	for _, num := range nums {
		uniq[num] = true
	}
	return uniq
}

func inOneNotInTwo(a, b map[int]bool) []int {
	var nums []int
	for k := range a {
		if !b[k] {
			nums = append(nums, k)
		}
	}
	return nums
}

// hash
// time complexity: O(n + m)
// space complexity: O(max(n, m))
func findDifference2(nums1 []int, nums2 []int) [][]int {
	m := make(map[int]int, len(nums1)+len(nums2))
	for _, num := range nums1 {
		m[num] = 1
	}
	for _, num := range nums2 {
		if val, ok := m[num]; ok && val != 2 {
			m[num] = 0
		} else {
			m[num] = 2
		}
	}

	var n1, n2 []int
	for k, v := range m {
		if v == 1 {
			n1 = append(n1, k)
		}
		if v == 2 {
			n2 = append(n2, k)
		}
	}
	return [][]int{n1, n2}
}

// use array
// time complexity: O(n + m)
// space complexity: O(max(n, m))
func findDifference3(nums1 []int, nums2 []int) [][]int {
	return [][]int{uniq(nums2, nums1), uniq(nums1, nums2)}
}

func uniq(s1, s2 []int) []int {
	d := [2001]bool{}
	for _, v1 := range s1 {
		d[v1+1000] = true
	}

	var res []int
	for _, v2 := range s2 {
		if !d[v2+1000] {
			res = append(res, v2)
			// prevent duplicates from being added
			d[v2+1000] = true
		}
	}
	return res
}
