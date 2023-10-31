package main

import (
	"strconv"
	"unicode"
)

func validWordAbbreviation(word, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if 'a' <= abbr[j] && abbr[j] <= 'z' && word[i] == abbr[j] {
			i, j = i+1, j+1
		} else if '0' <= abbr[j] && abbr[j] <= '9' {
			// leading 0
			if abbr[j] == '0' {
				return false
			}

			num := 0
			for j < len(abbr) && '0' <= abbr[j] && abbr[j] <= '9' {
				num = num*10 + int(abbr[j]-'0')
				j++
			}
			i += num
		} else {
			return false
		}
	}
	return i == len(word) && j == len(abbr)
}

func validWordAbbreviation2(word, abbr string) bool {
	n, i := 0, 0
	for _, r := range abbr {
		if unicode.IsLetter(r) {
			i += n
			if i >= len(word) || rune(word[i]) != r {
				return false
			}
			i, n = i+1, 0
		} else if unicode.IsNumber(r) {
			n = n*10 + int(r-'0')
			// leading 0
			if n == 0 {
				return false
			}
		} else {
			return false
		}
	}
	return len(word)-i == n
}

func validWordAbbreviation3(word, abbr string) bool {
	// safety conversion for indexing a char
	wordRune := []rune(word)
	abbrRune := []rune(abbr)

	wi, ai := 0, 0
	for wi < len(wordRune) && ai < len(abbrRune) {
		if unicode.IsLetter(abbrRune[ai]) {
			i, j := wi, ai
			for ai < len(abbrRune) && unicode.IsLetter(abbrRune[ai]) {
				wi, ai = wi+1, ai+1
			}

			if string(wordRune[i:wi]) != string(abbrRune[j:ai]) {
				return false
			}
		} else if unicode.IsNumber(abbrRune[ai]) {
			// leading 0
			if abbrRune[ai] == '0' {
				return false
			}

			j := ai
			for ai < len(abbrRune) && unicode.IsNumber(abbrRune[ai]) {
				ai++
			}

			n, _ := strconv.Atoi(string(abbrRune[j:ai]))
			wi += n
		} else {
			return false
		}
	}
	return wi == len(wordRune) && ai == len(abbrRune)
}
