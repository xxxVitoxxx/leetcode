package main

import "strings"

// use two pointers
func isSubsequence(s, t string) bool {
	var i int
	for j := 0; j < len(t); j++ {
		if i < len(s) && s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

// use two pointers
func isSubsequence2(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func isSubsequence3(s, t string) bool {
	for _, r := range s {
		if i := strings.IndexRune(t, r); i == -1 {
			return false
		} else {
			t = t[i+1:]
		}
	}
	return true
}
