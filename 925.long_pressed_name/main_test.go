package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLongPressedName(t *testing.T) {
	tests := []struct {
		name       string
		inputName  string
		inputTyped string
		output     bool
	}{
		{"example1", "alex", "aaleex", true},
		{"example2", "saeed", "ssaaedd", false},
		{"example3", "vicctoory", "vvvicctorrrry", false},
		{"example4", "mike", "rmmikkke", false},
		{"example5", "alex", "aaleexa", false},
		{"example6", "alex", "aaleelx", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isLongPressedName(tt.inputName, tt.inputTyped)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isLongPressedName2(tt.inputName, tt.inputTyped)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isLongPressedName3(tt.inputName, tt.inputTyped)
			assert.Equal(t, tt.output, res)
		})
	}
}
