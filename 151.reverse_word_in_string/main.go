package main

import (
	"sort"
	"strings"
)

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

func reverseWords2(s string) string {
	var strs []string
	pnt1, pnt2 := 0, 0
	for pnt1 < len(s) && pnt2 < len(s) {
		for pnt1 < len(s) && s[pnt1] == ' ' {
			pnt1++
		}
		pnt2 = pnt1
		for pnt2 < len(s) && s[pnt2] != ' ' {
			pnt2++
		}

		// 避免最後面都是空白
		if pnt1 < len(s) {
			strs = append(strs, s[pnt1:pnt2])
			pnt1 = pnt2 + 1
		}
	}

	left, right := 0, len(strs)-1
	for left < right {
		strs[left], strs[right] = strs[right], strs[left]
		left, right = left+1, right-1
	}

	return strings.Join(strs, " ")
}

func reverseWords3(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		j := i
		for j < len(s) && s[j] != ' ' {
			j++
		}

		if str == "" {
			str = s[i:j]
		} else {
			str = s[i:j] + " " + str
		}
		i = j
	}
	return str
}

func reverseWords4(s string) string {
	strs := strings.Fields(s)
	sort.SliceStable(strs, func(i, j int) bool {
		return i > j
	})
	return strings.Join(strs, " ")
}
