package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output int
	}{
		{"example1", []int{1, 1, 1}, 2},
		{"example2", []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{"example3", []int{1, 1, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestSubarray(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestSubarray2(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}
}
