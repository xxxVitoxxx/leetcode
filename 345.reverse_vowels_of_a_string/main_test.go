package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"example1", "hello", "holle"},
		{"example2", "leetcode", "leotcede"},
		{"example3", "AhPojiEtu", "uhPEjiotA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseVowels(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseVowels2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
