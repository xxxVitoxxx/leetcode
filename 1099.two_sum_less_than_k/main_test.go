package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumLessThanK(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output int
	}{
		{"example1", []int{34, 23, 1, 24, 75, 33, 54, 8}, 60, 58},
		{"example2", []int{10, 20, 30}, 15, -1},
		{"example3", []int{20, 23, 1, 40, 24, 75, 33, 54, 8}, 60, 57},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSumLessThanK(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSumLessThanK2(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
