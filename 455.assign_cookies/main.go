package main

import "sort"

// brute force
func findContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var res int
	for i := range g {
		for j := range s {
			if s[j] >= g[i] {
				res++
				s = append(s[:j], s[j+1:]...)
				break
			}
		}
	}
	return res
}

// use two pointers
func findContentChildren2(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i, j = i+1, j+1
		} else {
			j++
		}
	}
	return i
}
