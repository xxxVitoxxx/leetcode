package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{"example1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"example2", []int{1, 1}, 1},
		{"example3", []int{4, 3, 2, 1, 4}, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxArea(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxArea2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
