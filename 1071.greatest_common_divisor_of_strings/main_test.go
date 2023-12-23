package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name   string
		str1   string
		str2   string
		output string
	}{
		{"example1", "ABCABC", "ABC", "ABC"},
		{"example2", "ABABAB", "ABAB", "AB"},
		{"example3", "LEET", "CODE", ""},
		{"example4", "ABCDEF", "ABC", ""},
		{"example5", "TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXX"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := gcdOfStrings(tt.str1, tt.str2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := gcdOfStrings2(tt.str1, tt.str2)
			assert.Equal(t, tt.output, res)
		})
	}
}
