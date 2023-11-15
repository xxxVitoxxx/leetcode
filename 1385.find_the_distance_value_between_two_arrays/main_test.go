package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDistanceValue(t *testing.T) {
	tests := []struct {
		name   string
		arr1   []int
		arr2   []int
		d      int
		output int
	}{
		{"example1", []int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
		{"example2", []int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
		{"example3", []int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6, 1},
		{"example4", []int{-3, 10, 2, 8, 0, 10}, []int{-9, -1, -4, -9, -8}, 9, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTheDistanceValue(tt.arr1, tt.arr2, tt.d)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTheDistanceValue2(tt.arr1, tt.arr2, tt.d)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTheDistanceValue3(tt.arr1, tt.arr2, tt.d)
			assert.Equal(t, tt.output, res)
		})
	}
}
