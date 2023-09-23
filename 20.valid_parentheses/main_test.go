package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"example1", "()", true},
		{"example1", "()[]{}", true},
		{"example1", "(]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isValid(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
