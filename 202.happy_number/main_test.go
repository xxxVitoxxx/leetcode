package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{"example1", 19, true},
		{"example2", 2, false},
		{"example3", 125, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isHappy(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isHappy2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isHappy3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
