package main

import (
	"unicode"
)

// use two pointer
func reverseOnlyLetters(s string) string {
	sr := []rune(s)
	left, right := 0, len(sr)-1
	for left < right {
		for left < right && !unicode.IsLetter(sr[left]) {
			left++
		}
		for left < right && !unicode.IsLetter(sr[right]) {
			right--
		}
		sr[left], sr[right] = sr[right], sr[left]
		left, right = left+1, right-1
	}
	return string(sr)
}

func reverseOnlyLetters2(s string) string {
	sr := make([]rune, len(s))
	position := len(s) - 1
	for i, r := range s {
		if !unicode.IsLetter(r) {
			sr[i] = r
		} else {
			for !unicode.IsLetter(rune(s[position])) {
				position--
			}
			sr[i] = rune(s[position])
			position--
		}
	}
	return string(sr)
}
