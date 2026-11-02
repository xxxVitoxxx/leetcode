package main

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{{
		name:     "example1",
		text1:    "abcde",
		text2:    "ace",
		expected: 3,
	}, {
		name:     "example2",
		text1:    "abc",
		text2:    "abc",
		expected: 3,
	}, {
		name:     "example3",
		text1:    "abc",
		text2:    "def",
		expected: 0,
	}, {
		name:     "example4",
		text1:    "abcb",
		text2:    "acb",
		expected: 3,
	}, {
		name:     "example5",
		text1:    "abc",
		text2:    "acb",
		expected: 2,
	}, {
		name:     "example6",
		text1:    "bsbininm",
		text2:    "jmjkbkjkv",
		expected: 1,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestCommonSubsequence(tt.text1, tt.text2)
			if actual != tt.expected {
				t.Errorf("got %v, want: %v", actual, tt.expected)
			}
		})
	}
}
