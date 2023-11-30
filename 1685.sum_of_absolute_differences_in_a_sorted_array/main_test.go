package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumAbsoluteDifferences(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{2, 3, 5}, []int{4, 3, 5}},
		{"example2", []int{1, 4, 6, 8, 10}, []int{24, 15, 13, 15, 21}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getSumAbsoluteDifferences(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getSumAbsoluteDifferences2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getSumAbsoluteDifferences3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getSumAbsoluteDifferences4(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
