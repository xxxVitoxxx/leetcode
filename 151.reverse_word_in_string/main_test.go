package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		output string
	}{
		{"example1", "the sky is blue", "blue is sky the"},
		{"example2", "  hello world  ", "world hello"},
		{"example3", "a good   example", "example good a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseWords(tt.s)
			assert.Equal(t, tt.output, res)
		})
	}
}
