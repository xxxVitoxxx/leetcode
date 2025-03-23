package main

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{{
		name:     "example1",
		x:        2.0,
		n:        10,
		expected: 1024.0,
	}, {
		name:     "example2",
		x:        2.1,
		n:        3,
		expected: 9.261,
	}, {
		name:     "example3",
		x:        2.0,
		n:        -2,
		expected: 0.25,
	}, {
		name:     "example4",
		x:        5.0,
		n:        0,
		expected: 1.0,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := myPow(tt.x, tt.n)
			if math.Round(actual*1000)/1000 != tt.expected {
				t.Errorf("got: %f, want: %f", actual, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := myPow2(tt.x, tt.n)
			if math.Round(actual*1000)/1000 != tt.expected {
				t.Errorf("got: %f, want: %f", actual, tt.expected)
			}
		})
	}
}
