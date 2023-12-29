package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		k      int
		output int
	}{
		{"example1", "abciiidef", 3, 3},
		{"example2", "aeiou", 2, 2},
		{"example3", "leetcode", 3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxVowels(tt.s, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
