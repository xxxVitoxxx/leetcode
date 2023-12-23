package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		output       []bool
	}{
		{"example1", []int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{"example2", []int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{"example3", []int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := kidsWithCandies(tt.candies, tt.extraCandies)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
