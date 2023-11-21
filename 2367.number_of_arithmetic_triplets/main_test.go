package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticTriplets(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		diff   int
		output int
	}{
		{"example1", []int{0, 1, 4, 6, 7, 10}, 3, 2},
		{"example2", []int{4, 5, 6, 7, 8, 9}, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arithmeticTriplets(tt.nums, tt.diff)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arithmeticTriplets2(tt.nums, tt.diff)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arithmeticTriplets3(tt.nums, tt.diff)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arithmeticTriplets4(tt.nums, tt.diff)
			assert.Equal(t, tt.output, res)
		})
	}
}
