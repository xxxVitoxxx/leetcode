package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output []string
	}{
		{"example1", []string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{"example2", []string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := commonChars(tt.input)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := commonChars2(tt.input)
			assert.ElementsMatch(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := commonChars3(tt.input)
			assert.ElementsMatch(t, tt.output, res)
		})
	}
}
