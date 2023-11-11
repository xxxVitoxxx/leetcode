package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortEvenOdd(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 1, 2, 3}, []int{2, 3, 4, 1}},
		{"example2", []int{2, 1}, []int{2, 1}},
		{"example3", []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}, []int{9, 46, 15, 45, 15, 41, 27, 34, 32, 31, 33, 31, 36, 26, 36, 16, 44, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortEvenOdd(tt.input)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestSortEvenOdd2(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 1, 2, 3}, []int{2, 3, 4, 1}},
		{"example2", []int{2, 1}, []int{2, 1}},
		{"example3", []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}, []int{9, 46, 15, 45, 15, 41, 27, 34, 32, 31, 33, 31, 36, 26, 36, 16, 44, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortEvenOdd2(tt.input)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestSortEvenOdd3(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 1, 2, 3}, []int{2, 3, 4, 1}},
		{"example2", []int{2, 1}, []int{2, 1}},
		{"example3", []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}, []int{9, 46, 15, 45, 15, 41, 27, 34, 32, 31, 33, 31, 36, 26, 36, 16, 44, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortEvenOdd3(tt.input)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestSortEvenOdd4(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 1, 2, 3}, []int{2, 3, 4, 1}},
		{"example2", []int{2, 1}, []int{2, 1}},
		{"example3", []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}, []int{9, 46, 15, 45, 15, 41, 27, 34, 32, 31, 33, 31, 36, 26, 36, 16, 44, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortEvenOdd4(tt.input)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
