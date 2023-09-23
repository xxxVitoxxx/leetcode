package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"example1", "A", 1},
		{"example2", "AB", 28},
		{"example3", "ZY", 701},
		{"example3", "ZZA", 18253},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := titleToNumber(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestTitleToNumber2(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"example1", "A", 1},
		{"example2", "AB", 28},
		{"example3", "ZY", 701},
		{"example3", "ZZA", 18253},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := titleToNumber2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
