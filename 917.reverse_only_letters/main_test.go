package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseOnlyLetters(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"example1", "ab-cd", "dc-ba"},
		{"example2", "a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
		{"example3", "Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},
		{"example4", "*^f&ka!f@@@&#(#po)", "*^o&pf!a@@@&#(#kf)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseOnlyLetters(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseOnlyLetters2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
