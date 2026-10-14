package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{{
		name:     "example1",
		s:        "babad",
		expected: []string{"bab", "aba"},
	}, {
		name:     "example2",
		s:        "cbbd",
		expected: []string{"bb"},
	}, {
		name:     "example3",
		s:        "a",
		expected: []string{"a"},
	}, {
		name:     "example4",
		s:        "ac",
		expected: []string{"a", "c"},
	}, {
		name:     "example5",
		s:        "ccc",
		expected: []string{"ccc"},
	}, {
		name:     "example6",
		s:        "rabAkopr",
		expected: []string{"r", "a", "b", "A", "k", "o", "p", "r"},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestPalindrome(tt.s)

			var anyEqual bool
			for i := range tt.expected {
				if actual == tt.expected[i] {
					anyEqual = true
					break
				}
			}
			if !anyEqual {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestPalindrome2(tt.s)

			var anyEqual bool
			for i := range tt.expected {
				if actual == tt.expected[i] {
					anyEqual = true
					break
				}
			}
			if !anyEqual {
				t.Errorf("got: %v, want: %v", actual, tt.expected)
			}
		})
	}
}
