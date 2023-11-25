package main

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var res [][]int
	for pnt1, pnt2 := 0, 0; pnt1 < len(nums1) || pnt2 < len(nums2); {
		if pnt1 == len(nums1) && pnt2 < len(nums2) {
			res = append(res, nums2[pnt2:]...)
			break
		}
		if pnt2 == len(nums2) && pnt1 < len(nums1) {
			res = append(res, nums1[pnt1:]...)
			break
		}

		if nums1[pnt1][0] == nums2[pnt2][0] {
			res = append(res, []int{nums1[pnt1][0], nums1[pnt1][1] + nums2[pnt2][1]})
			pnt1, pnt2 = pnt1+1, pnt2+1
		} else if nums1[pnt1][0] < nums2[pnt2][0] {
			res = append(res, nums1[pnt1])
			pnt1++
		} else {
			res = append(res, nums2[pnt2])
			pnt2++
		}
	}
	return res
}

func mergeArrays2(nums1 [][]int, nums2 [][]int) [][]int {
	res := make([][]int, len(nums1)+len(nums2))
	idx, pnt1, pnt2 := 0, 0, 0
	for pnt1 < len(nums1) && pnt2 < len(nums2) {
		if nums1[pnt1][0] == nums2[pnt2][0] {
			res[idx] = []int{nums1[pnt1][0], nums1[pnt1][1] + nums2[pnt2][1]}
			idx, pnt1, pnt2 = idx+1, pnt1+1, pnt2+1
		} else if nums1[pnt1][0] < nums2[pnt2][0] {
			res[idx] = nums1[pnt1]
			idx, pnt1 = idx+1, pnt1+1
		} else {
			res[idx] = nums2[pnt2]
			idx, pnt2 = idx+1, pnt2+1
		}
	}

	for ; pnt1 < len(nums1); pnt1++ {
		res[idx] = nums1[pnt1]
		idx++
	}
	for ; pnt2 < len(nums2); pnt2++ {
		res[idx] = nums2[pnt2]
		idx++
	}
	return res[:idx]
}
