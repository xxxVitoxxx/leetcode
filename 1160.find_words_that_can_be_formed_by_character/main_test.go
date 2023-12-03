package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		chars  string
		output int
	}{
		{"example1", []string{"cat", "bt", "hat", "tree"}, "atach", 6},
		{"example1", []string{"hello", "world", "leetcode"}, "welldonehoneyr", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countCharacters(tt.words, tt.chars)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countCharacters2(tt.words, tt.chars)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countCharacters3(tt.words, tt.chars)
			assert.Equal(t, tt.output, res)
		})
	}
}
