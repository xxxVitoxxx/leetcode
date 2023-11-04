package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output bool
	}{
		{"example1", []int{1, 2, 3, 1}, true},
		{"example2", []int{1, 2, 3, 4}, false},
		{"example3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsDuplicate(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsDuplicate2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsDuplicate3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
