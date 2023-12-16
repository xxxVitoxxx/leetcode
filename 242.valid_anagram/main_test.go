package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name   string
		inputS string
		inputT string
		output bool
	}{
		{"example1", "anagram", "nagaram", true},
		{"example2", "rat", "car", false},
		{"example2", "index", "idx", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram(tt.inputS, tt.inputT)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram2(tt.inputS, tt.inputT)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram3(tt.inputS, tt.inputT)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram4(tt.inputS, tt.inputT)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram5(tt.inputS, tt.inputT)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram6(tt.inputS, tt.inputT)
			assert.Equal(t, tt.output, res)
		})
	}
}
