package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  [][]int
	}{
		{"example1", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{"example2", 1, [][]int{{1}}},
		{"example1", 7, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate(%d) = %v, want %v", tt.input, got, tt.want)
			}

			got2 := generate2(tt.input)
			if !reflect.DeepEqual(got2, tt.want) {
				t.Errorf("generate2(%d) = %v, want %v", tt.input, got2, tt.want)
			}
		})
	}
}
