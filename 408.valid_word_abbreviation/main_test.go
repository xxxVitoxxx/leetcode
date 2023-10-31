package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidWordAbbreviation(t *testing.T) {
	fmt.Println(len("internationalization"))
	tests := []struct {
		name   string
		word   string
		abbr   string
		output bool
	}{
		{"example1", "internationalization", "i12iz4n", true},
		{"example2", "apple", "a2e", false},
		{"example3", "apple", "1ppl1", true},
		{"example4", "internationalization", "i5a11o1", true},
		{"example5", "abbreviation", "abb0reviation", false},
		{"example6", "hi", "2i", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := validWordAbbreviation(tt.word, tt.abbr)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := validWordAbbreviation2(tt.word, tt.abbr)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := validWordAbbreviation3(tt.word, tt.abbr)
			assert.Equal(t, tt.output, res)
		})
	}
}
