package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheArrayConcVal(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output int64
	}{
		{"example1", []int{7, 52, 2, 4}, 596},
		{"example2", []int{5, 14, 13, 8, 12}, 673},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTheArrayConcVal(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTheArrayConcVal2(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findTheArrayConcVal3(tt.nums)
			assert.Equal(t, tt.output, res)
		})
	}
}
