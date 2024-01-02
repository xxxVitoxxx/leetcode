package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		output string
	}{
		{"example1", "3[a]2[bc]", "aaabcbc"},
		{"example2", "3[a2[c]]", "accaccacc"},
		{"example3", "2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"example4", "abc3[cd]xyz", "abccdcdcdxyz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := decodeString(tt.s)
			assert.Equal(t, tt.output, res)
		})
	}
}
