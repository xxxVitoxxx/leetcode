package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"example1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"example2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"example3", []int{0, 0}, []int{0, 0}},
		{"example4", []int{0, 4, 0}, []int{0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := productExceptSelf(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := productExceptSelf2(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
