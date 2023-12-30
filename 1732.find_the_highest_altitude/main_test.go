package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name   string
		gain   []int
		output int
	}{
		{"example1", []int{-5, 1, 5, 0, -7}, 1},
		{"example2", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := largestAltitude(tt.gain)
			assert.Equal(t, tt.output, res)
		})
	}
}
