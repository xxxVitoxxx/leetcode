package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStrobogrammatic(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"example1", "69", true},
		{"example2", "6", false},
		{"example3", "88", true},
		{"example4", "962", false},
		{"example5", "541", false},
		{"example6", "8601098", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isStrobogrammatic(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isStrobogrammatic2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isStrobogrammatic3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
