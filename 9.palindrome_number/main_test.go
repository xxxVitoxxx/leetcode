package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{"example1", 121, true},
		{"example2", -121, false},
		{"example3", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
