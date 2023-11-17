package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxK(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output int
	}{
		{"example1", []int{-1, 2, -3, 3}, 3},
		{"example2", []int{-1, 10, 6, 7, -7, 1}, 7},
		{"example3", []int{-10, 8, 6, 7, -2, -3}, -1},
		{"example4", []int{-9, -43, 24, -23, -16, -30, -38, -30}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxK(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxK2(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxK3(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxK4(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMaxK5(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}
}
