package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalPoints(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output int
	}{
		{"example1", []string{"5", "2", "C", "D", "+"}, 30},
		{"example2", []string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
		{"example3", []string{"1", "C"}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := calPoints(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
