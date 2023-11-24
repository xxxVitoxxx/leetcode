package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{"example1", [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, 4},
		{"example2", [][]int{{1}, {1}}, 0},
		{"example3", [][]int{{1, 4}, {0, 5}}, 4},
		{"example4", [][]int{{1, 5}, {3, 4}}, 3},
		{"example5", [][]int{{-3, -3}, {-3, -2}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxDistance(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
