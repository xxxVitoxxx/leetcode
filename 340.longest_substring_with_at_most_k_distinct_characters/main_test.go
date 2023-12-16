package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		k      int
		output int
	}{
		{"example1", "eceba", 2, 3},
		{"example2", "aa", 1, 2},
		{"example3", "ababcbcbaaabbdef", 2, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringKDistinct(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringKDistinct2(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringKDistinct3(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringKDistinct4(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
