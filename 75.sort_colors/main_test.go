package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result []int
	}{
		{"example1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"example2", []int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.result, tt.nums))
		})
	}
}

func TestSortColors2(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result []int
	}{
		{"example1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"example2", []int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors2(tt.nums)
			assert.True(t, reflect.DeepEqual(tt.result, tt.nums))
		})
	}
}
