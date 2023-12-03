package main

import "strings"

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	var ss []string
	for _, str := range strs {
		if str == "" {
			continue
		}
		ss = append(ss, str)
	}

	left, right := 0, len(ss)-1
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left, right = left+1, right-1
	}
	return strings.Join(ss, " ")
}
