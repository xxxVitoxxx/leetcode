package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		output string
	}{
		{"example1", "abc", "pqr", "apbqcr"},
		{"example2", "ab", "pqrs", "apbqrs"},
		{"example3", "abcd", "pq", "apbqcd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeAlternately(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeAlternately2(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeAlternately3(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeAlternately4(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := mergeAlternately5(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}
}
