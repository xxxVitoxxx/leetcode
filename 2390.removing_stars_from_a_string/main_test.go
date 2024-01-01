package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		output string
	}{
		{"example1", "leet**cod*e", "lecoe"},
		{"example1", "erase*****", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeStars(tt.s)
			assert.Equal(t, tt.output, res)
		})
	}
}
