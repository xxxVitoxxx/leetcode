package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output int
	}{
		{"example1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"example2", []int{1, 2, 3}, -1},
		{"example3", []int{2, 1, -1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pivotIndex(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := pivotIndex2(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}
}
