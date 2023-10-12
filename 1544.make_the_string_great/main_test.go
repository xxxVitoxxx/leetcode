package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGood(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"example1", "leEeetcode", "leetcode"},
		{"example2", "abBAcC", ""},
		{"example3", "s", "s"},
		{"example4", "Pp", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := makeGood(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := makeGood2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := makeGood3(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
