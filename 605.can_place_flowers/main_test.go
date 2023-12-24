package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		output    bool
	}{
		{"example1", []int{1, 0, 0, 0, 1}, 1, true},
		{"example2", []int{1, 0, 0, 0, 1}, 2, false},
		{"example3", []int{1, 0, 0, 0, 0, 1}, 2, false},
		{"example4", []int{0, 0, 1, 0, 1}, 1, true},
		{"example5", []int{1}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flowerbed := make([]int, len(tt.flowerbed))
			copy(flowerbed, tt.flowerbed)
			res := canPlaceFlowers(flowerbed, tt.n)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flowerbed := make([]int, len(tt.flowerbed))
			copy(flowerbed, tt.flowerbed)
			res := canPlaceFlowers2(flowerbed, tt.n)
			assert.Equal(t, tt.output, res)
		})
	}
}
