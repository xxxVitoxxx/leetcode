package main

import (
	"reflect"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {
	tests := []struct {
		name    string
		spells  []int
		potions []int
		success int64
		output  []int
	}{
		{"example1", []int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7, []int{4, 0, 3}},
		{"example2", []int{3, 1, 2}, []int{8, 5, 8}, 16, []int{2, 0, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := successfulPairs(tt.spells, tt.potions, tt.success)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
