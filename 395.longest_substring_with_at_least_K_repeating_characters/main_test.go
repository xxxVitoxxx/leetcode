package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		k      int
		output int
	}{
		{"example1", "aaabb", 3, 3},
		{"example2", "ababbc", 2, 5},
		{"example3", "ababacb", 3, 0},
		{"example4", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", 2, 52},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestSubstring(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestSubstring2(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestSubstring3(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestSubstring4(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
