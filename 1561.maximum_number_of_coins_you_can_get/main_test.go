package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		name   string
		piles  []int
		output int
	}{
		{"example1", []int{2, 4, 1, 2, 7, 8}, 9},
		{"example2", []int{2, 4, 5}, 4},
		{"example3", []int{9, 8, 7, 6, 5, 1, 2, 3, 4}, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxCoins(tt.piles)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxCoins2(tt.piles)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxCoins3(tt.piles)
			assert.Equal(t, tt.output, res)
		})
	}
}
