package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"example1", "00110011", 6},
		{"example2", "10101", 4},
		{"example3", "000111000", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBinarySubstrings(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBinarySubstrings2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBinarySubstrings3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBinarySubstrings4(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBinarySubstrings5(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
