package main

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		output []int
	}{
		{"example1", 2, []int{0, 1, 1}},
		{"example2", 5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBits(tt.n)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBits2(tt.n)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBits3(tt.n)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := countBits4(tt.n)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got %v want %v", res, tt.output)
			}
		})
	}
}
