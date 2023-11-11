package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParityII(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
		{"example2", []int{2, 3}, []int{2, 3}},
		{"example3", []int{4, 1, 1, 0, 1, 0}, []int{4, 1, 0, 1, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := distinguishEvenOddByVal(tt.output)
			res := sortArrayByParityII(tt.input)
			resEven, resOdd := distinguishEvenOddByIdx(res)
			assert.ElementsMatch(t, even, resEven)
			assert.ElementsMatch(t, odd, resOdd)

		})
	}
}

func TestSortArrayByParityII2(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
		{"example2", []int{2, 3}, []int{2, 3}},
		{"example3", []int{4, 1, 1, 0, 1, 0}, []int{4, 1, 0, 1, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := distinguishEvenOddByVal(tt.output)
			res := sortArrayByParityII2(tt.input)
			resEven, resOdd := distinguishEvenOddByIdx(res)
			assert.ElementsMatch(t, even, resEven)
			assert.ElementsMatch(t, odd, resOdd)

		})
	}
}

func TestSortArrayByParityII3(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
		{"example2", []int{2, 3}, []int{2, 3}},
		{"example3", []int{4, 1, 1, 0, 1, 0}, []int{4, 1, 0, 1, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := distinguishEvenOddByVal(tt.output)
			res := sortArrayByParityII3(tt.input)
			resEven, resOdd := distinguishEvenOddByIdx(res)
			assert.ElementsMatch(t, even, resEven)
			assert.ElementsMatch(t, odd, resOdd)

		})
	}
}

func TestSortArrayByParityII4(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"example1", []int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
		{"example2", []int{2, 3}, []int{2, 3}},
		{"example3", []int{4, 1, 1, 0, 1, 0}, []int{4, 1, 0, 1, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := distinguishEvenOddByVal(tt.output)
			res := sortArrayByParityII4(tt.input)
			resEven, resOdd := distinguishEvenOddByIdx(res)
			assert.ElementsMatch(t, even, resEven)
			assert.ElementsMatch(t, odd, resOdd)

		})
	}
}

func distinguishEvenOddByVal(nums []int) (even, odd []int) {
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return
}

func distinguishEvenOddByIdx(nums []int) (even, odd []int) {
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	return
}
