package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestToChar(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		c      byte
		output []int
	}{
		{"example1", "loveleetcode", 'e', []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"example2", "abaa", 'b', []int{1, 0, 1, 2}},
		{"example3", "aaab", 'b', []int{3, 2, 1, 0}},
		{"example4", "abcdefga", 'a', []int{0, 1, 2, 3, 3, 2, 1, 0}},
		{"example5", "uiiiiuiiii", 'u', []int{0, 1, 2, 2, 1, 0, 1, 2, 3, 4}},
		{"example6", "aaaaaacaaacacccaaaacacaaaaaaac", 'c', []int{6, 5, 4, 3, 2, 1, 0, 1, 2, 1, 0, 1, 0, 0, 0, 1, 2, 2, 1, 0, 1, 0, 1, 2, 3, 4, 3, 2, 1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := shortestToChar(tt.s, tt.c)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := shortestToChar2(tt.s, tt.c)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := shortestToChar3(tt.s, tt.c)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := shortestToChar4(tt.s, tt.c)
			assert.ElementsMatch(t, tt.output, res)
		})
	}
}
