package main

import "golang.org/x/exp/slices"

// sort and binary search
// time complexity: O((m+n)*logn) where m is length of the spells and n is length of the potions
// sort use pdsort in the worst case, it can be O(n * logn). foreach spells and binary search can be O(m*logn).
// so result is O((m+n)*logn)
// space complexity: O(m)
func successfulPairs(spells, potions []int, success int64) []int {
	slices.Sort(potions)
	res := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		left, right := 0, len(potions)-1
		for left <= right {
			mid := (left + right) / 2
			if int64(spells[i]*potions[mid]) < success {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		res[i] = len(potions) - left
	}
	return res
}
