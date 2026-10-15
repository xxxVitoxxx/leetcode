package main

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{{
		name:     "example1",
		s:        "leetcode",
		wordDict: []string{"leet", "code"},
		expected: true,
	}, {
		name:     "example2",
		s:        "applepenapple",
		wordDict: []string{"apple", "pen"},
		expected: true,
	}, {
		name:     "example3",
		s:        "catsandog",
		wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		expected: false,
	}, {
		name:     "example4",
		s:        "catsanddog",
		wordDict: []string{"and", "dog", "cats", "cat"},
		expected: true,
	}, {
		name:     "example5",
		s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		expected: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := wordBreak(tt.s, tt.wordDict)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := wordBreak2(tt.s, tt.wordDict)
			if actual != tt.expected {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
