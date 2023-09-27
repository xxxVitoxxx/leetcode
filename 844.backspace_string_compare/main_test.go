package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		name   string
		inputS string
		inputT string
		output bool
	}{
		{"example1", "ab#c", "ad#c", true},
		{"example2", "ab##", "c#d#", true},
		{"example3", "a#c", "b", false},
		{"example4", "#cv", "ca#d#v", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := backspaceCompare(tt.inputS, tt.inputT)
			assert.Equal(t, tt.output, res)
		})
	}
}
