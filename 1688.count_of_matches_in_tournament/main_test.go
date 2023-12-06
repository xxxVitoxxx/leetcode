package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfMatches(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"example1", 7, 6},
		{"example2", 14, 13},
		{"example3", 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := numberOfMatches(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := numberOfMatches2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
