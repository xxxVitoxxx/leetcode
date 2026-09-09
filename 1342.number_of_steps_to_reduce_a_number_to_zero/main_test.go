package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfStep1(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"example1", 14, 6},
		{"example2", 8, 4},
		{"example3", 123, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, numberOfSteps1(tt.input))
		})

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, numberOfSteps2(tt.input))
		})
	}
}
