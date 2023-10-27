package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfusingNumberII(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"example1", 20, 6},
		{"example2", 100, 19},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := confusingNumberII(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := confusingNumberII2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
