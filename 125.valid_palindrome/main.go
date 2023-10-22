package main

import (
	"strings"
	"unicode"
)

// brute force
func isPalindrome(s string) bool {
	var buf []byte
	for i := range s {
		if 'a' <= s[i] && s[i] <= 'z' || '0' <= s[i] && s[i] <= '9' {
			buf = append(buf, s[i])
		} else if 'A' <= s[i] && s[i] <= 'Z' {
			buf = append(buf, s[i]+32)
		}
	}

	// two pointers
	left, right := 0, len(buf)-1
	for left < right {
		if buf[left] != buf[right] {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}

// use two pointers
func isPalindrome2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !checkAlphanumeric(s[left]) {
			left++
		}

		for left < right && !checkAlphanumeric(s[right]) {
			right--
		}

		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}

func checkAlphanumeric(b byte) bool {
	if 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z' ||
		'0' <= b && b <= '9' {
		return true
	}
	return false
}

func isPalindrome3(s string) bool {
	transformed := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}

	clearStr := strings.Map(transformed, s)
	left, right := 0, len(clearStr)-1
	for left < right {
		if clearStr[left] != clearStr[right] {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}
