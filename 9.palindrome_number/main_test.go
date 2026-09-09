package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{"example1", 121, true},
		{"example2", -121, false},
		{"example3", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.output {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.input, got, tt.output)
			}

			got2 := isPalindrome2(tt.input)
			if got2 != tt.output {
				t.Errorf("isPalindrome2(%d) = %v, want %v", tt.input, got2, tt.output)
			}

		})
	}
}
