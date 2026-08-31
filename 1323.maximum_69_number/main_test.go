package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"example1", 9669, 9969},
		{"example2", 9999, 9999},
		{"example3", 96, 99},
		{"example4", 9996, 9999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximum69Number(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
