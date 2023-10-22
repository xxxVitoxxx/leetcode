package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"example1", "A man, a plan, a canal: Panama", true},
		{"example2", "race a car", false},
		{"example3", " ", true},
		{"example4", "0P", false},
		{"example5", "--IaM0aMA-?0mA)))i", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
