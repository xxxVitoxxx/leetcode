package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output int
	}{
		{"example1", []int{10, 5, 2, 6}, 100, 8},
		{"example1", []int{1, 2, 3}, 0, 0},
		{"example1", []int{1, 1, 1, 1, 1}, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := numSubarrayProductLessThanK(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
