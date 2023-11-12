package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiStringMatch(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []int
	}{
		{"example1", "IDID", []int{0, 4, 1, 3, 2}},
		{"example2", "III", []int{0, 1, 2, 3}},
		{"example3", "DDI", []int{3, 2, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := diStringMatch(tt.input)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := diStringMatch2(tt.input)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := diStringMatch3(tt.input)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
