package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output float64
	}{
		{"example1", []int{1, 12, -5, -6, 50, 3}, 4, float64(12.75)},
		{"example2", []int{5}, 1, float64(5)},
		{"example3", []int{-1}, 1, float64(-1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxAverage(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxAverage2(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
