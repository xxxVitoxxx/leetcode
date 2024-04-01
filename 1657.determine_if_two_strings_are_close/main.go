package main

import (
	"reflect"

	"golang.org/x/exp/slices"
)

// hash
// time complexity: O(n)
// space complexity: O(1) as the maximum size of each hashmap would be 26
// we are using constant extra space.
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1Map := make(map[int]int)
	word2Map := make(map[int]int)
	for i := 0; i < len(word1); i++ {
		word1Map[int(word1[i]-'a')]++
		word2Map[int(word2[i]-'a')]++
	}

	var word1KeyList, word1FrequencyList []int
	for k, v := range word1Map {
		word1KeyList = append(word1KeyList, k)
		word1FrequencyList = append(word1FrequencyList, v)
	}

	var word2KeyList, word2FrequencyList []int
	for k, v := range word2Map {
		word2KeyList = append(word2KeyList, k)
		word2FrequencyList = append(word2FrequencyList, v)
	}

	slices.Sort(word1KeyList)
	slices.Sort(word2KeyList)
	if !reflect.DeepEqual(word1KeyList, word2KeyList) {
		return false
	}

	slices.Sort(word1FrequencyList)
	slices.Sort(word2FrequencyList)
	return reflect.DeepEqual(word1FrequencyList, word2FrequencyList)
}

// using frequency array map
// time complexity: O(n)
// space complexity: O(1) as we use constant extra space of size 26 to store the frequency map.
func closeStrings2(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := [26]int{}
	freq2 := [26]int{}
	for i := 0; i < len(word1); i++ {
		freq1[word1[i]-'a']++
		freq2[word2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if freq1[i] > 0 && freq2[i] == 0 || freq2[i] > 0 && freq1[i] == 0 {
			return false
		}
	}

	slices.Sort(freq1[:])
	slices.Sort(freq2[:])

	return freq1 == freq2
}

// using bitwise operation and frequency array map
// time complexity: O(n)
// space complexity: O(1)
func closeStrings3(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var freq1, freq2 [26]int
	var bit1, bit2 int
	for i := 0; i < len(word1); i++ {
		freq1[word1[i]-'a']++
		freq2[word2[i]-'a']++
		bit1 = bit1 | (1 << (word1[i] - 'a'))
		bit2 = bit2 | (1 << (word2[i] - 'a'))
	}

	// check the characters of two words are the same
	if bit1 != bit2 {
		return false
	}

	slices.Sort(freq1[:])
	slices.Sort(freq2[:])
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
