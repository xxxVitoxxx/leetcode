package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output int
	}{
		{"example1", []int{1, 2, 3, 4}, 5, 2},
		{"example2", []int{3, 1, 3, 4, 3}, 6, 1},
		{"example3", []int{3, 5, 1, 5}, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxOperations(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxOperations2(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
