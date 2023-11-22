package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"example1", []int{3, 2, 1}, []int{1, 2, 3}},
		{"example1", []int{1, 1, 5}, []int{1, 5, 1}},
		{"example1", []int{1, 3, 2}, []int{2, 1, 3}},
		{"example1", []int{1, 5, 1}, []int{5, 1, 1}},
		{"example1", []int{2, 3, 1, 3, 3}, []int{2, 3, 3, 1, 3}},
		{"example1", []int{2, 1, 2, 2, 2, 2, 2, 1}, []int{2, 2, 1, 1, 2, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, tt.nums))
		})
	}
}

func TestNextPermutation2(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"example1", []int{3, 2, 1}, []int{1, 2, 3}},
		{"example1", []int{1, 1, 5}, []int{1, 5, 1}},
		{"example1", []int{1, 3, 2}, []int{2, 1, 3}},
		{"example1", []int{1, 5, 1}, []int{5, 1, 1}},
		{"example1", []int{2, 3, 1, 3, 3}, []int{2, 3, 3, 1, 3}},
		{"example1", []int{2, 1, 2, 2, 2, 2, 2, 1}, []int{2, 2, 1, 1, 2, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation2(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, tt.nums))
		})
	}
}

func TestNextPermutation3(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"example1", []int{3, 2, 1}, []int{1, 2, 3}},
		{"example1", []int{1, 1, 5}, []int{1, 5, 1}},
		{"example1", []int{1, 3, 2}, []int{2, 1, 3}},
		{"example1", []int{1, 5, 1}, []int{5, 1, 1}},
		{"example1", []int{2, 3, 1, 3, 3}, []int{2, 3, 3, 1, 3}},
		{"example1", []int{2, 1, 2, 2, 2, 2, 2, 1}, []int{2, 2, 1, 1, 2, 2, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation3(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, tt.nums))
		})
	}
}
