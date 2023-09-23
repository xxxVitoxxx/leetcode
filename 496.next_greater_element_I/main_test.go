package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		nums2  []int
		output []int
	}{
		{"example1", []int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{"example2", []int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		{"example3", []int{1, 3, 5, 2, 6}, []int{6, 5, 8, 3, 2, 1, 7}, []int{7, 7, 8, 7, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := nextGreaterElement(tt.nums1, tt.nums2)
			assert.Equal(t, tt.output, res)
		})
	}
}
