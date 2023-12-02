package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStringsAreEqual(t *testing.T) {
	tests := []struct {
		name   string
		word1  []string
		word2  []string
		output bool
	}{
		{"example1", []string{}, []string{}, true},
		{"example2", []string{"ab", "c"}, []string{"a", "bc"}, true},
		{"example3", []string{"a", "cb"}, []string{"ab", "c"}, false},
		{"example4", []string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
		{"example5", []string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{"example6", []string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{"example7", []string{"ecxarwyyy", "ppf", "tdyayjd"}, []string{"ecxarwyyyppft", "dyayj", "q"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arrayStringsAreEqual(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arrayStringsAreEqual2(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arrayStringsAreEqual3(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := arrayStringsAreEqual4(tt.word1, tt.word2)
			assert.Equal(t, tt.output, res)
		})
	}
}
