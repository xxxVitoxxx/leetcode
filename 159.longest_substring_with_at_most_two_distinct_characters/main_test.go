package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"example1", "eceba", 3},
		{"example2", "ccaabbb", 5},
		{"example3", "abcabcabc", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringTwoDistinct(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstringTwoDistinct2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
