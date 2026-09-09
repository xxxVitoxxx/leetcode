package main

import (
	"testing"
)

func TestAddDigits(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"example1", 38, 2},
		{"example2", 0, 0},
		{"example3", 10, 1},
		{"example4", 872, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addDigits(tt.input)
			if got != tt.want {
				t.Errorf("addDigits(%d) = %d, want %d", tt.input, got, tt.want)
			}

			got2 := addDigits2(tt.input)
			if got2 != tt.want {
				t.Errorf("addDigits2(%d) = %d, want %d", tt.input, got2, tt.want)
			}
		})
	}
}
