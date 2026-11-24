package main

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{{
		name:     "example1",
		nums:     []int{1, 2, 2},
		expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
