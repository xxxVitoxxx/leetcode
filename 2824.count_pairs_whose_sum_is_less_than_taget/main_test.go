package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPairs(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		output int
	}{
		{"example1", []int{-1, 1, 2, 3, 1}, 2, 3},
		{"example1", []int{-6, 2, 5, -2, -7, -1, 3}, -2, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countPairs(tt.nums, tt.target)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countPairs2(tt.nums, tt.target)
			assert.Equal(t, tt.output, res)
		})
	}
}
