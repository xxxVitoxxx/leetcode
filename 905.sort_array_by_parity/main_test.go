package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{1, 4, 2, 0}, []int{4, 2, 0, 1}},
		{"example2", []int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
		{"example3", []int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := distinguishEvenOdd(tt.input)
			res := sortArrayByParity(tt.input)
			assert.ElementsMatch(t, even, res[:len(even)])
			assert.ElementsMatch(t, odd, res[len(even):])
		})
	}
}

func TestSortArrayByParity2(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{1, 4, 2, 0}, []int{4, 2, 0, 1}},
		{"example2", []int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
		{"example3", []int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := distinguishEvenOdd(tt.input)
			res := sortArrayByParity2(tt.input)
			assert.ElementsMatch(t, even, res[:len(even)])
			assert.ElementsMatch(t, odd, res[len(even):])
		})
	}
}

func TestSortArrayByParity3(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{1, 4, 2, 0}, []int{4, 2, 0, 1}},
		{"example2", []int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
		{"example3", []int{0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := distinguishEvenOdd(tt.input)
			res := sortArrayByParity3(tt.input)
			assert.ElementsMatch(t, even, res[:len(even)])
			assert.ElementsMatch(t, odd, res[len(even):])
		})
	}
}

func distinguishEvenOdd(nums []int) (even, odd []int) {
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return
}
