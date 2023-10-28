package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		res   []int
	}{
		{"example1", []int{0, 1, 3, 0, 12}, []int{1, 3, 12, 0, 0}},
		{"example2", []int{0}, []int{0}},
		{"example3", []int{0, -4, -2, 0, 2, 6}, []int{-4, -2, 2, 6, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.input)
			assert.True(t, reflect.DeepEqual(tt.input, tt.res))
		})
	}
}

func TestMoveZeroes2(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		res   []int
	}{
		{"example1", []int{0, 1, 3, 0, 12}, []int{1, 3, 12, 0, 0}},
		{"example2", []int{0}, []int{0}},
		{"example3", []int{0, -4, -2, 0, 2, 6}, []int{-4, -2, 2, 6, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes2(tt.input)
			assert.True(t, reflect.DeepEqual(tt.input, tt.res))
		})
	}
}

func TestMoveZeroes3(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		res   []int
	}{
		{"example1", []int{0, 1, 3, 0, 12}, []int{1, 3, 12, 0, 0}},
		{"example2", []int{0}, []int{0}},
		{"example3", []int{0, -4, -2, 0, 2, 6}, []int{-4, -2, 2, 6, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes3(tt.input)
			assert.True(t, reflect.DeepEqual(tt.input, tt.res))
		})
	}
}

func TestMoveZeroes4(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		res   []int
	}{
		{"example1", []int{0, 1, 3, 0, 12}, []int{1, 3, 12, 0, 0}},
		{"example2", []int{0}, []int{0}},
		{"example3", []int{0, -4, -2, 0, 2, 6}, []int{-4, -2, 2, 6, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes4(tt.input)
			assert.True(t, reflect.DeepEqual(tt.input, tt.res))
		})
	}
}
