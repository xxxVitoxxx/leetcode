package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateZeros(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		output []int
	}{
		{"example1", []int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{"example2", []int{1, 2, 3}, []int{1, 2, 3}},
		{"example3", []int{0, 1, 0, 2, 0}, []int{0, 0, 1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.arr)
			assert.True(t, reflect.DeepEqual(tt.output, tt.arr))
		})
	}
}

func TestDuplicateZeros2(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		output []int
	}{
		{"example1", []int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{"example2", []int{1, 2, 3}, []int{1, 2, 3}},
		{"example3", []int{0, 1, 0, 2, 0}, []int{0, 0, 1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros2(tt.arr)
			assert.True(t, reflect.DeepEqual(tt.output, tt.arr))
		})
	}
}

func TestDuplicateZeros3(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		output []int
	}{
		{"example1", []int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{"example2", []int{1, 2, 3}, []int{1, 2, 3}},
		{"example3", []int{0, 1, 0, 2, 0}, []int{0, 0, 1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros3(tt.arr)
			assert.True(t, reflect.DeepEqual(tt.output, tt.arr))
		})
	}
}

func TestDuplicateZeros4(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		output []int
	}{
		{"example1", []int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{"example2", []int{1, 2, 3}, []int{1, 2, 3}},
		{"example3", []int{0, 1, 0, 2, 0}, []int{0, 0, 1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros4(tt.arr)
			assert.True(t, reflect.DeepEqual(tt.output, tt.arr))
		})
	}
}
