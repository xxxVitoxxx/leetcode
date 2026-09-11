package main

import "testing"

func TestIsUgly(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"example1", 8, true},
		{"example2", 6, true},
		{"example3", 14, false},
		{"example4", 1, true},
		{"example5", 0, false},
		{"example6", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isUgly(tt.input)
			if got != tt.want {
				t.Errorf("isUgly(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
