package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStr(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		k      int
		output string
	}{
		{"example1", "abcdefg", 2, "bacdfeg"},
		{"example2", "abcd", 2, "bacd"},
		{"example3", "a", 2, "a"},
		{"example4", "abcdefg", 1, "abcdefg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseStr(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseStr2(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseStr3(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseStr4(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
