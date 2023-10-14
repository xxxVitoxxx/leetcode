package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularArrayLoop(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output bool
	}{
		{"example1", []int{2, -1, 1, 2, 2}, true},
		{"example2", []int{-1, -2, -3, -4, -5, 6}, false},
		{"example3", []int{1, -1, 5, 1, 4}, true},
		{"example4", []int{-2, 1, -1, -2, -2}, false},
		{"example5", []int{6, -1, 5, 11, 4}, true},
		{"example6", []int{-8, -1, 5}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := circularArrayLoop(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := circularArrayLoop2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
