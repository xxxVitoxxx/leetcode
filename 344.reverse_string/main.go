package main

import "sort"

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left, right = left+1, right-1
	}
}

func reverseString2(s []byte) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func reverseString3(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
