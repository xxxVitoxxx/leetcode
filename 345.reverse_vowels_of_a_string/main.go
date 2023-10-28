package main

import "strings"

func reverseVowels(s string) string {
	buf := []byte(s)
	left, right := 0, len(buf)-1
	for left < right {
		for left < right && !isVowels(buf[left]) {
			left++
		}
		for left < right && !isVowels(buf[right]) {
			right--
		}
		buf[left], buf[right] = buf[right], buf[left]
		left, right = left+1, right-1
	}
	return string(buf)
}

func isVowels(b byte) bool {
	return b == 'a' || b == 'A' ||
		b == 'e' || b == 'E' ||
		b == 'i' || b == 'I' ||
		b == 'o' || b == 'O' ||
		b == 'u' || b == 'U'
}

func reverseVowels2(s string) string {
	buf := []byte(s)
	vowels := "aAeEiIoOuU"
	left, right := 0, len(buf)-1
	for left < right {
		for left < right && !strings.Contains(vowels, string(buf[left])) {
			left++
		}

		for left < right && !strings.Contains(vowels, string(buf[right])) {
			right--
		}

		buf[left], buf[right] = buf[right], buf[left]
		left, right = left+1, right-1
	}
	return string(buf)
}
