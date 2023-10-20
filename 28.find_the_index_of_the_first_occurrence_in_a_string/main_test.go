package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		output   int
	}{
		{"example1", "sadbutsad", "sad", 0},
		{"example2", "leetcode", "leeto", -1},
		{"example3", "ibighaveabigdream", "bigdr", 9},
		{"example4", "aaa", "aaaa", -1},
		{"example5", "mississippi", "issipi", -1},
		{"example6", "a", "a", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := strStr(tt.haystack, tt.needle)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := strStr2(tt.haystack, tt.needle)
			assert.Equal(t, tt.output, res)
		})
	}
}
