package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		val    int
		output int
	}{
		{"example1", []int{3, 2, 2, 3}, 3, 2},
		{"example2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{"example3", []int{1, 4, 2, 1, 1, 3, 1}, 1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeElement(tt.input, tt.val)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestRemoveElement2(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		val    int
		output int
	}{
		{"example1", []int{3, 2, 2, 3}, 3, 2},
		{"example2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{"example3", []int{1, 4, 2, 1, 1, 3, 1}, 1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeElement2(tt.input, tt.val)
			assert.Equal(t, tt.output, res)
		})
	}
}
