package main

import (
	"strings"
)

func reverseWords(s string) string {
	sr := []rune(s)
	var i int
	for j, r := range sr {
		if r == ' ' {
			reverse(sr[i:j])
			i = j + 1
		}
		if j == len(sr)-1 {
			reverse(sr[i:])
		}
	}
	return string(sr)
}

func reverse(r []rune) []rune {
	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
	return r
}

func reverseWords2(s string) string {
	// sl := strings.Fields(s)
	sl := strings.Split(s, " ")
	for i, word := range sl {
		wr := []rune(word)
		for l, r := 0, len(wr)-1; l < r; l, r = l+1, r-1 {
			wr[l], wr[r] = wr[r], wr[l]
		}
		sl[i] = string(wr)
	}
	return strings.Join(sl, " ")
}

func reverseWords3(s string) string {
	var sw strings.Builder
	var i int
	sr := []rune(s)
	for j, r := range sr {
		if r == ' ' {
			for back := j - 1; back >= i; back-- {
				sw.WriteRune(sr[back])
			}
			sw.WriteRune(' ')
			i = j + 1
		}
		if j == len(sr)-1 {
			for back := j; back >= i; back-- {
				sw.WriteRune(sr[back])
			}
		}
	}
	return sw.String()
}

func reverseWords4(s string) string {
	var res string
	var word string
	for _, char := range strings.Split(s, "") {
		if char != " " {
			word = char + word
		} else {
			res += word + char
			word = ""
		}
	}
	return res + word
}
