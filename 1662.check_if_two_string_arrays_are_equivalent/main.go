package main

import "strings"

// brute force
func arrayStringsAreEqual(word1, word2 []string) bool {
	w1, w2 := "", ""
	for _, word := range word1 {
		w1 += word
	}

	for _, word := range word2 {
		w2 += word
	}
	return w1 == w2
}

func arrayStringsAreEqual2(word1, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

// use two pointer
func arrayStringsAreEqual3(word1, word2 []string) bool {
	w1pnt, s1pnt := 0, 0
	w2pnt, s2pnt := 0, 0
	for w1pnt < len(word1) && w2pnt < len(word2) {
		if word1[w1pnt][s1pnt] != word2[w2pnt][s2pnt] {
			return false
		}

		s1pnt, s2pnt = s1pnt+1, s2pnt+1

		if s1pnt == len(word1[w1pnt]) {
			w1pnt, s1pnt = w1pnt+1, 0
		}

		if s2pnt == len(word2[w2pnt]) {
			w2pnt, s2pnt = w2pnt+1, 0
		}
	}
	return w1pnt == len(word1) && w2pnt == len(word2)
}

// use djb2 hash function
func arrayStringsAreEqual4(word1, word2 []string) bool {
	return hashWord(word1) == hashWord(word2)
}

// djb2: http://www.cse.yorku.ca/~oz/hash.html
// https://theartincode.stanis.me/008-djb2/
// https://www.jianshu.com/p/7f2a34288892
func hashWord(words []string) int {
	hash := 5381
	for _, word := range words {
		for _, r := range word {
			// hash = hash<<5 + hash + int(r)
			hash = hash*33 ^ int(r)
		}
	}
	return hash
}
