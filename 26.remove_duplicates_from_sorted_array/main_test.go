package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDeplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int

		// input[:output]
		expected []int
		output   int
	}{
		{"example1", []int{1, 1, 2}, []int{1, 2}, 2},
		{"example2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			length := removeDuplicates(tt.input)
			assert.Equal(t, tt.output, length)
			assert.ElementsMatch(t, tt.expected, tt.input[:length])
		})
	}
}
