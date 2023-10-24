package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfusingNumber(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{"example1", 6, true},
		{"example2", 89, true},
		{"example3", 11, false},
		{"example4", 8000, true},
		{"example5", 6891689, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := confusingNumber(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := confusingNumber2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
