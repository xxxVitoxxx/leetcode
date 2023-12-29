package main

import (
	"strconv"
)

// brute force
// time complexity: O(n)
// space complexity: O(1)
func compress(chars []byte) int {
	count, pnt := 1, 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			count++
		} else {
			if count != 1 {
				num := strconv.Itoa(count)
				for l := 0; l < len(num); l++ {
					chars[pnt] = num[l]
					pnt++
				}
			}
			chars[pnt] = chars[i]
			pnt++
			count = 1
		}
		if i == len(chars)-1 {
			if count != 1 {
				num := strconv.Itoa(count)
				for l := 0; l < len(num); l++ {
					chars[pnt] = num[l]
					pnt++
				}
			}
		}
	}
	return pnt
}

// optimize compress
// time complexity: O(n)
// space complexity: O(1)
func compress2(chars []byte) int {
	var ans int
	for i := 0; i <= len(chars)-1; i++ {
		groupLen := 1
		for i < len(chars)-1 && chars[i] == chars[i+1] {
			groupLen++
			i++
		}
		chars[ans] = chars[i]
		ans++

		if groupLen > 1 {
			str := strconv.Itoa(groupLen)
			for l := 0; l < len(str); l++ {
				chars[ans] = str[l]
				ans++
			}
		}
	}
	return ans
}
