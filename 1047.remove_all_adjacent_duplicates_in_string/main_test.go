package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"example1", "abbaca", "ca"},
		{"example2", "azxxzy", "ay"},
		{"example3", "cbbac", "cac"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeDuplicates(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeDuplicates2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
