package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePrefix(t *testing.T) {
	tests := []struct {
		name   string
		word   string
		ch     byte
		output string
	}{
		{"example1", "abcdefd", 'd', "dcbaefd"},
		{"example2", "xyxzxe", 'z', "zxyxxe"},
		{"example3", "abcd", 'z', "abcd"},
		{"example3", "j", 'j', "j"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reversePrefix(tt.word, tt.ch)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reversePrefix2(tt.word, tt.ch)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reversePrefix3(tt.word, tt.ch)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reversePrefix4(tt.word, tt.ch)
			assert.Equal(t, tt.output, res)
		})
	}
}
