package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k    int
		n    int
		want [][]int
	}{
		{
			name: "example1",
			k:    3,
			n:    7,
			want: [][]int{
				{1, 2, 4},
			},
		},
		{
			name: "example2",
			k:    3,
			n:    9,
			want: [][]int{
				{1, 2, 6},
				{1, 3, 5},
				{2, 3, 4},
			},
		},
		{
			name: "example3",
			k:    4,
			n:    1,
			want: [][]int{},
		},
		{
			name: "example4",
			k:    4,
			n:    24,
			want: [][]int{
				{1, 6, 8, 9},
				{2, 5, 8, 9},
				{2, 6, 7, 9},
				{3, 4, 8, 9},
				{3, 5, 7, 9},
				{3, 6, 7, 8},
				{4, 5, 6, 9},
				{4, 5, 7, 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.k, tt.n)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
