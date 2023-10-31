package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		t      string
		output bool
	}{
		{"example1", "abc", "abcde", true},
		{"example2", "aec", "abcde", false},
		{"example3", "axc", "ahbgdc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSubsequence(tt.s, tt.t)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSubsequence2(tt.s, tt.t)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSubsequence3(tt.s, tt.t)
			assert.Equal(t, tt.output, res)
		})
	}
}
