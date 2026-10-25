package main

import (
	"sort"
	"strings"
)

// N is the length of strs, and K is the maximum length of a string is strs.
// time complexity: O(NK log K)
// space complexity: O(NK)
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		groups[key] = append(groups[key], str)
	}

	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func sortString(str string) string {
	sorted := strings.Split(str, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

// Categorize by Count
// time complexity: O(NK)
// space complexity: O(NK)
func groupAnagrams2(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, str := range strs {
		key := getKey(str)
		groups[key] = append(groups[key], str)
	}

	var result [][]string
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func getKey(str string) [26]int {
	var result [26]int
	for _, s := range str {
		result[s-'a']++
	}
	return result
}
