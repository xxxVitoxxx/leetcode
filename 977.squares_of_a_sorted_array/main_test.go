package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"example2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortedSquares(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestSortedSquares2(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"example2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortedSquares2(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestSortedSquares3(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"example2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortedSquares3(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestSortedSquares4(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"example2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortedSquares4(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
