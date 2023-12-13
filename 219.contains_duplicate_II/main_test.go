package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output bool
	}{
		{"example1", []int{1, 2, 3, 1}, 3, true},
		{"example2", []int{1, 0, 1, 1}, 1, true},
		{"example3", []int{1, 2, 3, 1, 2, 3}, 2, false},
		{"example3", []int{0, 1, 2, 3, 4, 0, 0, 7, 8, 9, 10, 11, 12, 0}, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsNearByDuplicate(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsNearByDuplicate2(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsNearByDuplicate3(tt.nums, tt.k)
			assert.Equal(t, tt.output, res)
		})
	}
}
