package main

import (
	"strings"
)

func reversePrefix(word string, ch byte) string {
	res := make([]byte, len(word))
	var reverse bool
	for i := range word {
		res[i] = word[i]
		if word[i] == ch && !reverse {
			reverse = true
			left, right := 0, i
			for left < right {
				res[left], res[right] = word[right], word[left]
				left, right = left+1, right-1
			}
		}
	}
	return string(res)
}

func reversePrefix2(word string, ch byte) string {
	var occ int
	for ; occ < len(word); occ++ {
		if word[occ] == ch {
			break
		}
	}

	if occ == len(word) {
		return word
	}

	res := make([]byte, occ+1)
	left, right := 0, occ
	for ; left <= right; left, right = left+1, right-1 {
		res[left], res[right] = word[right], word[left]
	}
	return string(res) + word[occ+1:]
}

func reversePrefix3(word string, ch byte) string {
	occ := strings.IndexByte(word, ch)
	if occ == -1 {
		return word
	}

	str := make([]byte, occ+1)
	left, right := 0, occ
	for left <= right {
		str[left], str[right] = word[right], word[left]
		left, right = left+1, right-1
	}
	return string(str) + word[occ+1:]
}

func reversePrefix4(word string, ch byte) string {
	res := []byte(word)
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			left, right := 0, i
			for left < right {
				res[left], res[right] = res[right], res[left]
				left, right = left+1, right-1
			}
			return string(res)
		}
	}
	return word
}
