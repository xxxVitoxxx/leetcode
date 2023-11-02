package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"example1", "Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"example2", "God Ding", "doG gniD"},
		{"example3", "Have a good day", "evaH a doog yad"},
		{"example3", "We'll be back", "ll'eW eb kcab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseWords(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseWords2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseWords3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseWords4(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
